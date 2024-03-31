package constants

type LanguageMap map[string]map[string]string

const (
	INTERNAL_SERVER_ERROR string = "internal_server_error"
)

var Languages = LanguageMap{
	"en": {
		INTERNAL_SERVER_ERROR: "Internal server error",
	},
	"id": {
		INTERNAL_SERVER_ERROR: "Terjadi kesalahan sistem",
	},
}
