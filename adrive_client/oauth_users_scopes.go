package alipan

import (
	"fmt"
	"github.com/niuhuan/alipan-go/adrive_client/protos"
	"github.com/niuhuan/alipan-go/common"
	"net/url"
)

func (c *AdriveClient) OauthUserScopes(_params *protos.OauthUserScopesParams) (*protos.OauthUserScopes, error) {
	apiUrl := fmt.Sprintf("%s/oauth/users/scopes", c.ApiHost)
	params := url.Values{}
	return common.DoGetQuery[protos.OauthUserScopes](c.AccessTokenLoader, c.Agent, apiUrl, params)
}
