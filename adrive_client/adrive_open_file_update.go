package alipan

import (
	"fmt"
	"github.com/niuhuan/alipan-go/adrive_client/protos"
	"github.com/niuhuan/alipan-go/common"
)

func (c *AdriveClient) AdriveOpenFileUpdate(params *protos.AdriveOpenFileUpdateParams) (*protos.AdriveOpenFile, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/update", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFile](c.AccessTokenLoader, c.Agent, apiUrl, params)
}
