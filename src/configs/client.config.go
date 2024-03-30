package configs

import (
	"os"
)

type S_ClientConfig struct {
	ArthaUrl string
}

func ClientConfig() S_ClientConfig {
	return S_ClientConfig{
		ArthaUrl: os.Getenv("ARTHA_URL"),
	}
}
