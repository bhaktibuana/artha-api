package constants

type LanguageMap map[string]map[string]string

const (
	INTERNAL_SERVER_ERROR string = "internal_server_error"
	URL_NOT_FOUND         string = "url_not_found"
	EMAIL_EXISTS          string = "email_exists"
	LOGIN_SUCCESS         string = "login_success"
	WRONG_MAIL_PASS       string = "wrong_mail_pass"
	UNVERIFIED_MAIL       string = "unverified_mail"
	REGISTER_SUCCESS      string = "register_success"
	REGISTER_FAILED       string = "register_failed"
	REQUEST_SUCCESS       string = "request_success"
	DATA_NOT_FOUND        string = "data_not_found"
	INVALID_USER          string = "invalid_user"
	INVALID_OTP           string = "invalid_otp"
	SET_2FA_FAILED        string = "set_2fa_failed"
	WRONG_PASS            string = "wrong_pass"
	INVALID_TOKEN         string = "invalid_token"
	MAIL_ALREADY_VERIFIED string = "mail_already_verified"
	VERIFY_MAIL_FAILED    string = "verify_mail_failed"
)

var Languages = LanguageMap{
	"en": {
		INTERNAL_SERVER_ERROR: "Internal server error.",
		URL_NOT_FOUND:         "URL not found.",
		EMAIL_EXISTS:          "Email already exists.",
		LOGIN_SUCCESS:         "Login success.",
		WRONG_MAIL_PASS:       "Wrong email or password.",
		UNVERIFIED_MAIL:       "The email has not been verified yet.",
		REGISTER_SUCCESS:      "Register success.",
		REGISTER_FAILED:       "Register failed.",
		REQUEST_SUCCESS:       "Request successful.",
		DATA_NOT_FOUND:        "Data not found.",
		INVALID_USER:          "Invalid user.",
		INVALID_OTP:           "Invalid OTP.",
		SET_2FA_FAILED:        "Set Two Factor Authentication failed.",
		WRONG_PASS:            "Incorrect password.",
		INVALID_TOKEN:         "Invalid token.",
		MAIL_ALREADY_VERIFIED: "Email already verified.",
		VERIFY_MAIL_FAILED:    "Failed to verify email.",
	},
	"id": {
		INTERNAL_SERVER_ERROR: "Terjadi kesalahan sistem.",
		URL_NOT_FOUND:         "URL tidak ditemukan.",
		EMAIL_EXISTS:          "Email sudah tersedia.",
		LOGIN_SUCCESS:         "Berhasil masuk.",
		WRONG_MAIL_PASS:       "Email atau kata sandi salah.",
		UNVERIFIED_MAIL:       "Email belum terverifikasi.",
		REGISTER_SUCCESS:      "Berhasil mendaftar.",
		REGISTER_FAILED:       "Gagal mendaftar.",
		REQUEST_SUCCESS:       "Permintaan berhasil.",
		DATA_NOT_FOUND:        "Data tidak ditemukan.",
		INVALID_USER:          "Pengguna tidak valid.",
		INVALID_OTP:           "OTP tidak valid.",
		SET_2FA_FAILED:        "Gagal Mengatur Otentikasi Dua Faktor.",
		WRONG_PASS:            "Kata sandi salah.",
		INVALID_TOKEN:         "Token tidak valid.",
		MAIL_ALREADY_VERIFIED: "Email sudah terverifikasi.",
		VERIFY_MAIL_FAILED:    "Verifikasi email gagal.",
	},
}
