package alipan

import (
	"fmt"
	"github.com/niuhuan/alipan-go/adrive_client/protos"
	"github.com/niuhuan/alipan-go/common"
)

func (c *AdriveClient) AdriveOpenFileGetDownloadUrl(params *protos.AdriveOpenFileGetDownloadUrlParams) (*protos.AdriveOpenFileGetDownloadUrl, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/getDownloadUrl", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileGetDownloadUrl](c.AccessTokenLoader, c.Agent, apiUrl, params)
}
