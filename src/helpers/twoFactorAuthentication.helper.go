package helpers

import (
	"artha-api/src/constants"
	"time"

	"github.com/xlzd/gotp"
)

// GenerateTOTPWithSecret Helper
/*
 * @param email string
 * @returns (string, string) (secretKey, uri)
 */
func GenerateTOTPWithSecret(email string) (secretKey string, uri string) {
	secretKey = gotp.RandomSecret(16)
	totp := gotp.NewDefaultTOTP(secretKey)
	uri = totp.ProvisioningUri(email, constants.APP_TOTP_ISSUE)
	return secretKey, uri
}

// VerifyOTP Helper
/*
 * @param otp string
 * @param secretKey string
 * @returns bool
 */
func VerifyOTP(otp, secretKey string) bool {
	totp := gotp.NewDefaultTOTP(secretKey)
	return totp.Verify(otp, time.Now().Unix())
}
