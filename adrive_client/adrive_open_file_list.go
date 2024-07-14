package alipan

import (
	"fmt"
	"github.com/niuhuan/alipan-go/adrive_client/protos"
	"github.com/niuhuan/alipan-go/common"
)

func (c *AdriveClient) AdriveOpenFileList(params *protos.AdriveOpenFileListParams) (*protos.AdriveOpenFileList, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/list", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileList](c.AccessTokenLoader, c.Agent, apiUrl, params)
}
