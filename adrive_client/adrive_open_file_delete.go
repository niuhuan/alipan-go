package alipan

import (
	"fmt"
	"github.com/niuhuan/alipan-go/adrive_client/protos"
	"github.com/niuhuan/alipan-go/common"
)

func (c *AdriveClient) AdriveOpenFileDelete(params *protos.AdriveOpenFileDeleteParams) (*protos.AdriveOpenFileDelete, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/recyclebin/trash", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileDelete](c.AccessTokenLoader, c.Agent, apiUrl, params)
}
