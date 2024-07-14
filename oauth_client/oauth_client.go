package oauth_client

import (
	"net/http"
)

type OauthClient struct {
	Agent        *http.Client
	ApiHost      string
	ClientId     string
	ClientSecret string
}

func NewOauthClient(clientId string, clientSecret string) *OauthClient {
	return &OauthClient{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		ApiHost:      "https://openapi.alipan.com",
		Agent:        &http.Client{},
	}
}
