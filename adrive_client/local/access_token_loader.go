package local

import (
	"errors"
	"github.com/niuhuan/alipan-go/common"
	"github.com/niuhuan/alipan-go/oauth_client"
	"github.com/niuhuan/alipan-go/oauth_client/protos"
)

type AccessToken struct {
	AccessToken  string `json:"access_token" toml:"access_token"`
	ExpiresIn    int64  `json:"expires_in" toml:"expires_in"`
	RefreshToken string `json:"refresh_token" toml:"refresh_token"`
	TokenType    string `json:"token_type" toml:"token_type"`
	CreatedAt    int64  `json:"created_at" toml:"created_at"`
}

func FromOauthAccessToken(token *protos.OauthAccessToken) AccessToken {
	return AccessToken{
		AccessToken:  token.AccessToken,
		ExpiresIn:    token.ExpiresIn,
		RefreshToken: token.RefreshToken,
		TokenType:    token.TokenType,
		CreatedAt:    common.Timestamp(),
	}
}

type AccessTokenStore interface {
	SaveAccessToken(token *AccessToken) error
	LoadAccessToken() (*AccessToken, error)
}

type OauthClientTokenLoader struct {
	AccessTokenStore AccessTokenStore
	OauthClient      *oauth_client.OauthClient
}

// LoadAccessToken impl AccessTokenLoad for OauthClientTokenLoader
func (o *OauthClientTokenLoader) LoadAccessToken() (*string, error) {
	a, e := o.AccessTokenStore.LoadAccessToken()
	if e != nil {
		return nil, e
	}
	if a == nil {
		return nil, errors.New("access token not found")
	}
	now := common.Timestamp()
	if now < a.CreatedAt+a.ExpiresIn*int64(3)/int64(4) {
		return &a.AccessToken, nil
	}
	token, e := o.OauthClient.OauthAccessToken(&protos.OauthAccessTokenParams{
		GrantType:    "refresh_token",
		RefreshToken: a.RefreshToken,
	})
	if e != nil {
		return nil, e
	}
	accessToken := FromOauthAccessToken(token)
	e = o.AccessTokenStore.SaveAccessToken(&accessToken)
	if e != nil {
		return nil, e
	}
	return &accessToken.AccessToken, nil
}
