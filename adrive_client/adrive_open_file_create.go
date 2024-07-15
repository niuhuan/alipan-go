package alipan

import (
	"fmt"
	"github.com/niuhuan/alipan-go/adrive_client/protos"
	"github.com/niuhuan/alipan-go/common"
)

func (c *AdriveClient) AdriveOpenFileCreate(params *protos.AdriveOpenFileCreateParams) (*protos.AdriveOpenFileCreate, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/create", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileCreate](c.AccessTokenLoader, c.Agent, apiUrl, params)
}
