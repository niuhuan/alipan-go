package alipan

import (
	"fmt"
	"github.com/niuhuan/alipan-go/adrive_client/protos"
	"github.com/niuhuan/alipan-go/common"
)

func (c *AdriveClient) AdriveOpenFileMove(params *protos.AdriveOpenFileMoveParams) (*protos.AdriveOpenFileMove, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/copy", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileMove](c.AccessTokenLoader, c.Agent, apiUrl, params)
}
