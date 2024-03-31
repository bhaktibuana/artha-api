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
	},
}
