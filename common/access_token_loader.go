package common

type AccessTokenLoader interface {
	LoadAccessToken() (*string, error)
}
