package configs

import (
	"os"
)

type S_SMTPConfig struct {
	SMTP_URL string
}

func SMTPConfig() S_SMTPConfig {
	return S_SMTPConfig{
		SMTP_URL: os.Getenv("ARTHA_SMTP_URL"),
	}
}
