package authService

import (
	"artha-api/src/apis"
	"artha-api/src/configs"
	"artha-api/src/constants"
	"artha-api/src/database"
	"artha-api/src/helpers"
	"artha-api/src/models"
	authRequest "artha-api/src/requests/auth"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Login Service
/*
 * @param context *gin.Context
 * @param request *authRequest.S_LoginRequest
 * @returns *models.Users
 */
func Login(context *gin.Context, request *authRequest.S_LoginRequest) *models.Users {
	var user models.Users

	if !request.Encrypted {
		request.Password = helpers.HashPassword(request.Password)
	}

	normalizedEmail := strings.ToLower(request.Email)

	filter := bson.M{"email": normalizedEmail, "password": request.Password}

	if err := database.Users.FindOne(context, filter).Decode(&user); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			helpers.HttpResponse(constants.WRONG_MAIL_PASS, http.StatusBadRequest, context, nil)
			return nil
		default:
			helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, nil)
			return nil
		}
	}

	if !user.DeletedAt.IsZero() {
		helpers.HttpResponse(constants.WRONG_MAIL_PASS, http.StatusBadRequest, context, nil)
		return nil
	}

	if user.Status == models.USER_STATUS_UNVERIFIED {
		helpers.HttpResponse(constants.UNVERIFIED_MAIL, http.StatusBadRequest, context, nil)
		return nil
	}

	return &user
}

// Register Service
/*
 * @param context *gin.Context
 * @param request *authRequest.S_RegisterRequest
 * @returns *models.Users
 */
func Register(context *gin.Context, request *authRequest.S_RegisterRequest) *models.Users {
	var user models.Users
	var username, tag, accountId string

	if !request.Encrypted {
		request.Password = helpers.HashPassword(request.Password)
	}

	for {
		username, tag, accountId = helpers.GenerateAccountId()

		filter := bson.M{"username": username, "tag": tag}

		if err := database.Users.FindOne(context, filter).Decode(&user); err != nil {
			if err == mongo.ErrNoDocuments {
				break
			}
		}
	}

	normalizedEmail := strings.ToLower(request.Email)

	user = models.Users{
		ID:          primitive.NewObjectID(),
		Name:        request.Name,
		Username:    username,
		Tag:         tag,
		Email:       normalizedEmail,
		Password:    request.Password,
		AccountType: models.USER_ACCOUNT_TYPE_USER,
		Status:      models.USER_STATUS_UNVERIFIED,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if _, err := database.Users.InsertOne(context, &user); err != nil {
		helpers.HttpResponse(constants.REGISTER_FAILED, http.StatusBadRequest, context, nil)
		return nil
	}

	claims := jwt.MapClaims{
		"id": user.ID,
	}

	token, _ := helpers.GenerateJWT(claims, 0)

	payload := map[string]string{
		"name":         user.Name,
		"username":     accountId,
		"email":        user.Email,
		"app_logo_url": configs.AppConfig().BASE_URL + "/artha-logo.png",
		"login_url":    configs.ClientConfig().ARTHA_URL + "/login",
		"action_url":   configs.ClientConfig().ARTHA_URL + "/verifyAccount?token=" + token,
	}

	apis.APIArthaSMTP().MailVerification(payload)

	return &user
}

// Me Service
/*
 * @param context *gin.Context
 * @returns *models.Users
 */
func Me(context *gin.Context) *models.Users {
	var user models.Users

	id, err := helpers.GetSelfID(context)
	if err != nil {
		helpers.HttpResponse(constants.INVALID_USER, http.StatusBadRequest, context, nil)
		return nil
	}

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	if err := database.Users.FindOne(context, filter).Decode(&user); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			helpers.HttpResponse(constants.DATA_NOT_FOUND, http.StatusNotFound, context, nil)
			return nil
		default:
			helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, nil)
			return nil
		}
	}

	if user.Status == models.USER_STATUS_UNVERIFIED {
		helpers.HttpResponse(constants.UNVERIFIED_MAIL, http.StatusBadRequest, context, nil)
		return nil
	}

	return &user
}

// verify2FAOTP Service
/*
 * @param context *gin.Context
 * @param id string
 * @param otp string
 * @returns *models.Users
 */
func verify2FAOTP(context *gin.Context, id, otp string) *models.Users {
	var user models.Users

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	if err := database.Users.FindOne(context, filter).Decode(&user); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			helpers.HttpResponse(constants.DATA_NOT_FOUND, http.StatusNotFound, context, nil)
			return nil
		default:
			helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, nil)
			return nil
		}
	}

	if !helpers.VerifyOTP(otp, user.Secret2FA) {
		helpers.HttpResponse(constants.INVALID_OTP, http.StatusBadRequest, context, nil)
		return nil
	}

	return &user
}

// LoginVerify2FA Service
/*
 * @param context *gin.Context
 * @param id string
 * @param request *authRequest.S_LoginVerify2FA
 * @returns *models.Users
 */
func LoginVerify2FA(context *gin.Context, id string, request *authRequest.S_LoginVerify2FA) *models.Users {
	user := verify2FAOTP(context, id, request.OTP)

	return user
}

// Verify2FA Service
/*
 * @param context *gin.Context
 * @param request *authRequest.S_Verify2FA
 * @returns *models.Users
 */
func Verify2FA(context *gin.Context, request *authRequest.S_Verify2FA) *models.Users {
	id, err := helpers.GetSelfID(context)
	if err != nil {
		helpers.HttpResponse(constants.INVALID_USER, http.StatusBadRequest, context, nil)
		return nil
	}

	user := verify2FAOTP(context, id, request.OTP)

	return user
}

type S_GetNew2FAUri struct {
	SecretKey string
	URI       string
}

// GetNew2FAUri Service
/*
 * @param context *gin.Context
 * @returns *models.Users
 */
func GetNew2FAUri(context *gin.Context) *S_GetNew2FAUri {
	var user models.Users

	id, err := helpers.GetSelfID(context)
	if err != nil {
		helpers.HttpResponse(constants.INVALID_USER, http.StatusBadRequest, context, nil)
		return nil
	}
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	if err := database.Users.FindOne(context, filter).Decode(&user); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			helpers.HttpResponse(constants.DATA_NOT_FOUND, http.StatusNotFound, context, nil)
			return nil
		default:
			helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, nil)
			return nil
		}
	}

	if user.Status == models.USER_STATUS_UNVERIFIED {
		helpers.HttpResponse(constants.UNVERIFIED_MAIL, http.StatusBadRequest, context, nil)
		return nil
	}

	secretKey, uri := helpers.GenerateTOTPWithSecret(user.Email)

	return &S_GetNew2FAUri{SecretKey: secretKey, URI: uri}
}

// Update2FASecret Service
/*
 * @param context *gin.Context
 * @param request *authRequest.S_Update2FASecret
 * @returns *models.Users
 */
func Update2FASecret(context *gin.Context, request *authRequest.S_Update2FASecret) *models.Users {
	var user models.Users

	id, err := helpers.GetSelfID(context)
	if err != nil {
		helpers.HttpResponse(constants.INVALID_USER, http.StatusBadRequest, context, nil)
		return nil
	}
	_id, _ := primitive.ObjectIDFromHex(id)

	if !request.Encrypted {
		request.Password = helpers.HashPassword(request.Password)
	}

	filter := bson.M{"_id": _id, "password": request.Password}

	if err := database.Users.FindOne(context, filter).Decode(&user); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			helpers.HttpResponse(constants.WRONG_PASS, http.StatusBadRequest, context, nil)
			return nil
		default:
			helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, nil)
			return nil
		}
	}

	if !helpers.VerifyOTP(request.OTP, request.SecretKey) {
		helpers.HttpResponse(constants.INVALID_OTP, http.StatusBadRequest, context, nil)
		return nil
	}

	updatePayload := bson.M{
		"$set": bson.M{
			"secret_2fa": request.SecretKey,
		},
	}

	if _, err := database.Users.UpdateOne(context, filter, updatePayload); err != nil {
		helpers.HttpResponse(constants.SET_2FA_FAILED, http.StatusBadRequest, context, nil)
		return nil
	}

	return &user
}

// VerifyEmail Service
/*
 * @param context *gin.Context
 * @param request *authRequest.S_VerifyEmail
 * @returns *models.Users
 */
func VerifyEmail(context *gin.Context, request *authRequest.S_VerifyEmail) *models.Users {
	var user models.Users

	claims, err := helpers.VerifyJwt(request.Token)
	if err != nil {
		helpers.HttpResponse(constants.INVALID_TOKEN, http.StatusBadRequest, context, nil)
		return nil
	}

	id, ok := claims["id"].(string)
	if !ok {
		helpers.HttpResponse(constants.INVALID_TOKEN, http.StatusBadRequest, context, nil)
		return nil
	}

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	if err = database.Users.FindOne(context, filter).Decode(&user); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			helpers.HttpResponse(constants.INVALID_USER, http.StatusNotFound, context, nil)
			return nil
		default:
			helpers.HttpResponse(constants.INTERNAL_SERVER_ERROR, http.StatusInternalServerError, context, nil)
			return nil
		}
	}

	if user.Status == models.USER_STATUS_VERIFIED {
		helpers.HttpResponse(constants.MAIL_ALREADY_VERIFIED, http.StatusBadRequest, context, filter)
		return nil
	}

	updatePayload := bson.M{
		"$set": bson.M{
			"status": models.USER_STATUS_VERIFIED,
		},
	}

	if _, err = database.Users.UpdateOne(context, filter, updatePayload); err != nil {
		helpers.HttpResponse(constants.VERIFY_MAIL_FAILED, http.StatusBadRequest, context, nil)
		return nil
	}

	return &user
}
