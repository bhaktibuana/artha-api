package authResult

import authService "artha-api/src/services/auth"

type S_GetNew2FAUri struct {
	SecretKey string `json:"secret_key"`
	URI       string `json:"artha_id"`
}

// GetNew2FAUri Request
/*
 * @param user *authService.S_GetNew2FAUri
 * @returns S_GetNew2FAUri
 */
func GetNew2FAUri(payload *authService.S_GetNew2FAUri) S_GetNew2FAUri {
	return S_GetNew2FAUri{
		SecretKey: payload.SecretKey,
		URI:       payload.URI,
	}
}
