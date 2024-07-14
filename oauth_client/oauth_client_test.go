package oauth_client

import (
	"github.com/BurntSushi/toml"
	"github.com/niuhuan/alipan-go/oauth_client/protos"
	"os"
	"testing"
)

func oauthClient(t *testing.T) *OauthClient {
	buff, err := os.ReadFile("../tests_data/app_info.toml")
	if err != nil {
		t.Fatal(err)
	}
	appInfo := map[string]string{}
	err = toml.Unmarshal(buff, &appInfo)
	if err != nil {
		t.Fatal(err)
	}
	client := NewOauthClient(appInfo["client_id"], appInfo["client_secret"])
	return client
}

func TestOauthClient_OauthAuthorizeUrl(t *testing.T) {
	client := oauthClient(t)
	params := &protos.OauthAuthorizeUrlParams{
		RedirectUri:  "http://localhost:58080/oauth_authorize",
		Scope:        "user:base,file:all:read,file:all:write,album:shared:read",
		ResponseType: "code",
	}
	url, err := client.OauthAuthorizeUrl(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
}

func TestOauthClient_OauthAccessToken(t *testing.T) {
	client := oauthClient(t)
	code := "code"
	token, err := client.OauthAccessToken(&protos.OauthAccessTokenParams{
		GrantType: "authorization_code",
		Code:      &code,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}
