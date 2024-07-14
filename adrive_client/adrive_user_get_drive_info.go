package alipan

import (
	"fmt"
	"github.com/niuhuan/alipan-go/adrive_client/protos"
	"github.com/niuhuan/alipan-go/common"
)

func (c *AdriveClient) AdriveUserGetDriveInfo(_params *protos.AdriveUserGetDriveInfoParams) (*protos.AdriveUserGetDriveInfo, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/user/getDriveInfo", c.ApiHost)
	return common.DoPostNoBody[protos.AdriveUserGetDriveInfo](c.AccessTokenLoader, c.Agent, apiUrl)
}
