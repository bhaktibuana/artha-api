package helpers

import (
	"artha-api/src/constants"
	"time"

	"github.com/xlzd/gotp"
)

func GenerateTOTPWithSecret(email string) (string, string) {
	secretKey := gotp.RandomSecret(16)
	totp := gotp.NewDefaultTOTP(secretKey)
	uri := totp.ProvisioningUri(email, constants.APP_TOTP_ISSUE)
	return secretKey, uri
}

func VerifyOTP(otp, secretKey string) bool {
	totp := gotp.NewDefaultTOTP(secretKey)
	return totp.Verify(otp, time.Now().Unix())
}
