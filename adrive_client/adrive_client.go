package alipan

import (
	"github.com/niuhuan/alipan-go/common"
	"net/http"
)

type AdriveClient struct {
	Agent             *http.Client
	ApiHost           string
	ClientId          string
	AccessTokenLoader common.AccessTokenLoader
}

func NewAdriveClient(clientId string, accessTokenLoader common.AccessTokenLoader) *AdriveClient {
	return &AdriveClient{
		ClientId:          clientId,
		ApiHost:           "https://openapi.alipan.com",
		Agent:             &http.Client{},
		AccessTokenLoader: accessTokenLoader,
	}
}
