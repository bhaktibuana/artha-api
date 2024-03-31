package configs

import (
	"os"
)

type S_ClientConfig struct {
	ARTHA_URL string
}

func ClientConfig() S_ClientConfig {
	return S_ClientConfig{
		ARTHA_URL: os.Getenv("ARTHA_URL"),
	}
}
