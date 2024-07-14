package alipan

import (
	"fmt"
	"github.com/niuhuan/alipan-go/adrive_client/protos"
	"github.com/niuhuan/alipan-go/common"
	"net/url"
)

func (c *AdriveClient) OauthUserInfo(_params *protos.OauthUserInfoParams) (*protos.OauthUserInfo, error) {
	apiUrl := fmt.Sprintf("%s/oauth/users/info", c.ApiHost)
	params := url.Values{}
	return common.DoGetQuery[protos.OauthUserInfo](c.AccessTokenLoader, c.Agent, apiUrl, params)
}
