package alipan

import (
	"fmt"
	"github.com/niuhuan/alipan-go/adrive_client/protos"
	"github.com/niuhuan/alipan-go/common"
)

func (c *AdriveClient) AdriveOpenFileCopy(params *protos.AdriveOpenFileCopyParams) (*protos.AdriveOpenFileCopy, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/copy", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileCopy](c.AccessTokenLoader, c.Agent, apiUrl, params)
}
