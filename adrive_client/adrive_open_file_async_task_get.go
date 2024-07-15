package alipan

import (
	"fmt"
	"github.com/niuhuan/alipan-go/adrive_client/protos"
	"github.com/niuhuan/alipan-go/common"
)

func (c *AdriveClient) AdriveOpenFileAsyncTaskGet(params *protos.AdriveOpenFileAsyncTaskGetParams) (*protos.AdriveOpenFileAsyncTaskGet, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/async_task/get", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileAsyncTaskGet](c.AccessTokenLoader, c.Agent, apiUrl, params)
}
