package oauth_client

import (
	"fmt"
	"github.com/niuhuan/alipan-go/common"
	"github.com/niuhuan/alipan-go/oauth_client/protos"
	"net/url"
)

func (c *OauthClient) OauthAccessToken(params *protos.OauthAccessTokenParams) (*protos.OauthAccessToken, error) {
	apiUrl := fmt.Sprintf("%v/oauth/access_token", c.ApiHost)
	formValues := url.Values{}
	formValues.Set("client_id", c.ClientId)
	formValues.Set("client_secret", c.ClientSecret)
	formValues.Set("grant_type", params.GetGrantType())
	formValues.Set("code", params.GetCode())
	formValues.Set("refresh_token", params.GetRefreshToken())
	formValues.Set("code_verifier", params.GetCodeVerifier())
	return common.DoFormRequest[protos.OauthAccessToken](nil, c.Agent, apiUrl, formValues)
}
