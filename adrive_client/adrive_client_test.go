package alipan

import (
	"github.com/BurntSushi/toml"
	"github.com/niuhuan/alipan-go/adrive_client/local"
	"github.com/niuhuan/alipan-go/adrive_client/protos"
	"github.com/niuhuan/alipan-go/oauth_client"
	"os"
	"testing"
)

func oauthClient(t *testing.T) *oauth_client.OauthClient {
	buff, err := os.ReadFile("../tests_data/app_info.toml")
	if err != nil {
		t.Fatal(err)
	}
	appInfo := map[string]string{}
	err = toml.Unmarshal(buff, &appInfo)
	if err != nil {
		t.Fatal(err)
	}
	client := oauth_client.NewOauthClient(appInfo["client_id"], appInfo["client_secret"])
	return client
}

type TomlAccessTokenStore struct{}

func (t *TomlAccessTokenStore) SaveAccessToken(token *local.AccessToken) error {
	buff, err := toml.Marshal(token)
	if err != nil {
		return err
	}
	err = os.WriteFile("../tests_data/access_token.toml", buff, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (t *TomlAccessTokenStore) LoadAccessToken() (*local.AccessToken, error) {
	buff, err := os.ReadFile("../tests_data/access_token.toml")
	if err != nil {
		return nil, err
	}
	token := &local.AccessToken{}
	err = toml.Unmarshal(buff, token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func adriveClient(t *testing.T) *AdriveClient {
	buff, err := os.ReadFile("../tests_data/app_info.toml")
	if err != nil {
		t.Fatal(err)
	}
	appInfo := map[string]string{}
	err = toml.Unmarshal(buff, &appInfo)
	if err != nil {
		t.Fatal(err)
	}
	client := NewAdriveClient(appInfo["client_id"], &local.OauthClientTokenLoader{
		AccessTokenStore: &TomlAccessTokenStore{},
		OauthClient:      oauthClient(t),
	})
	return client
}

func TestAdriveClient_OauthUserInfo(t *testing.T) {
	client := adriveClient(t)
	info, err := client.OauthUserInfo(&protos.OauthUserInfoParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_OauthUserScopes(t *testing.T) {
	client := adriveClient(t)
	info, err := client.OauthUserScopes(&protos.OauthUserScopesParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_AdriveUserGetSpaceInfo(t *testing.T) {
	client := adriveClient(t)
	info, err := client.AdriveUserGetSpaceInfo(&protos.AdriveUserGetSpaceInfoParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_AdriveUserGetDriveInfo(t *testing.T) {
	client := adriveClient(t)
	info, err := client.AdriveUserGetDriveInfo(&protos.AdriveUserGetDriveInfoParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_AdriveOpenFileList(t *testing.T) {
	client := adriveClient(t)
	info, err := client.AdriveOpenFileList(&protos.AdriveOpenFileListParams{
		DriveId:      "drive_id",
		ParentFileId: "root",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_AdriveOpenFileGet(t *testing.T) {
	client := adriveClient(t)
	info, err := client.AdriveOpenFileGet(&protos.AdriveOpenFileGetParams{
		DriveId: "drive_id",
		FileId:  "file_id",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_AdriveOpenFileGetByPath(t *testing.T) {
	client := adriveClient(t)
	info, err := client.AdriveOpenFileGetByPath(&protos.AdriveOpenFileGetByPathParams{
		DriveId:  "drive_id",
		FilePath: "/file_path",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_AdriveOpenFileGetDownloadUrl(t *testing.T) {
	client := adriveClient(t)
	info, err := client.AdriveOpenFileGetDownloadUrl(&protos.AdriveOpenFileGetDownloadUrlParams{
		DriveId: "drive_id",
		FileId:  "file_id",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}
