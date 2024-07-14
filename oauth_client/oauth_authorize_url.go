package oauth_client

import (
	"fmt"
	"net/url"

	"github.com/niuhuan/alipan-go/oauth_client/protos"
)

func (c *OauthClient) OauthAuthorizeUrl(params *protos.OauthAuthorizeUrlParams) (string, error) {
	url, err := url.Parse(c.ApiHost + "/oauth/authorize")
	if err != nil {
		return "", err
	}
	q := url.Query()
	q.Set("client_id", c.ClientId)
	q.Set("response_type", params.GetResponseType())
	q.Set("redirect_uri", params.GetRedirectUri())
	q.Set("scope", params.GetScope())
	if params.State != nil {
		q.Set("state", *params.State)
	}
	if params.Relogin != nil {
		q.Set("relogin", fmt.Sprintf("%t", *params.Relogin))
	}
	url.RawQuery = q.Encode()
	return url.String(), nil
}
