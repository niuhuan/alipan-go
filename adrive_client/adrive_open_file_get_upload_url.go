package alipan

import (
	"fmt"
	"github.com/niuhuan/alipan-go/adrive_client/protos"
	"github.com/niuhuan/alipan-go/common"
)

func (c *AdriveClient) AdriveOpenFileGetUploadUrl(params *protos.AdriveOpenFileGetUploadUrlParams) (*protos.AdriveOpenFileGetUploadUrl, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/getUploadUrl", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileGetUploadUrl](c.AccessTokenLoader, c.Agent, apiUrl, params)
}
