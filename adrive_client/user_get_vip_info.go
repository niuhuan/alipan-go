package alipan

import (
	"fmt"
	"github.com/niuhuan/alipan-go/adrive_client/protos"
	"github.com/niuhuan/alipan-go/common"
	"net/url"
)

func (c *AdriveClient) UserGetVipInfo(_params *protos.UserGetVipInfoParams) (*protos.UserGetVipInfo, error) {
	apiUrl := fmt.Sprintf("%s/oauth/users/info", c.ApiHost)
	params := url.Values{}
	return common.DoGetQuery[protos.UserGetVipInfo](c.AccessTokenLoader, c.Agent, apiUrl, params)
}
