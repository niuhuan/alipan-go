package alipan

import (
	"fmt"
	"github.com/niuhuan/alipan-go/adrive_client/protos"
	"github.com/niuhuan/alipan-go/common"
)

func (c *AdriveClient) AdriveOpenFileRecyclebinTrash(params *protos.AdriveOpenFileRecyclebinTrashParams) (*protos.AdriveOpenFileRecyclebinTrash, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/recyclebin/trash", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileRecyclebinTrash](c.AccessTokenLoader, c.Agent, apiUrl, params)
}
