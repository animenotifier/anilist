package anilist

import (
	"errors"

	"github.com/parnurzeal/gorequest"
)

// AccessToken is the anilist access token
var AccessToken = ""

// APIKeyID ...
var APIKeyID = ""

// APIKeySecret ...
var APIKeySecret = ""

// AuthorizeResponse ...
type AuthorizeResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expires     int    `json:"expires"`
	ExpiresIn   int    `json:"expires_in"`
}

// Authorize ...
func Authorize() error {
	request := gorequest.New().Post("https://anilist.co/api/auth/access_token")

	request.QueryData.Add("grant_type", "client_credentials")
	request.QueryData.Add("client_id", APIKeyID)
	request.QueryData.Add("client_secret", APIKeySecret)

	authorization := &AuthorizeResponse{}
	_, _, errs := request.EndStruct(authorization)

	if len(errs) > 0 {
		return errs[0]
	}

	AccessToken = authorization.AccessToken

	if AccessToken == "" {
		return errors.New("Access token is empty")
	}

	return nil
}
