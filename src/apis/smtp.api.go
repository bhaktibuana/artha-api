package apis

import (
	"artha-api/src/configs"
	"artha-api/src/helpers"
	"fmt"
)

type S_APIArthaSMTP struct {
	client helpers.S_Client
}

func APIArthaSMTP() *S_APIArthaSMTP {
	baseUrl := fmt.Sprintf("%s%s", configs.SMTPConfig().SMTP_URL, "/api/smtp")
	client := helpers.NewClient(baseUrl, "")
	return &S_APIArthaSMTP{client: *client}
}

func (api *S_APIArthaSMTP) MailVerification(payload interface{}) ([]byte, error) {
	endpoint := "/mail-verification"
	return api.client.Post(endpoint, payload)
}
