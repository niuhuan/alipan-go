package alipan

import (
	"fmt"
	"github.com/niuhuan/alipan-go/adrive_client/protos"
	"github.com/niuhuan/alipan-go/common"
)

func (c *AdriveClient) AdriveOpenFileComplete(params *protos.AdriveOpenFileCompleteParams) (*protos.AdriveOpenFile, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/complete", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFile](c.AccessTokenLoader, c.Agent, apiUrl, params)
}
