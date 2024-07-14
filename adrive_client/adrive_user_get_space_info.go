package alipan

import (
	"fmt"
	"github.com/niuhuan/alipan-go/adrive_client/protos"
	"github.com/niuhuan/alipan-go/common"
)

func (c *AdriveClient) AdriveUserGetSpaceInfo(_params *protos.AdriveUserGetSpaceInfoParams) (*protos.AdriveUserGetSpaceInfo, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/user/getSpaceInfo", c.ApiHost)
	return common.DoPostNoBody[protos.AdriveUserGetSpaceInfo](c.AccessTokenLoader, c.Agent, apiUrl)
}
