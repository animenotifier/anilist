package anilist

import (
	"fmt"
	"net/http"

	"github.com/aerogo/http/client"
)

var headers = client.Headers{
	"User-Agent":   "notify.moe testing",
	"Content-Type": "application/json",
}

// Query queries the AniList GraphQL API.
func Query(body interface{}, target interface{}) error {
	response, err := client.Post("https://graphql.anilist.co").Headers(headers).BodyJSON(body).EndStruct(target)

	if err != nil {
		return err
	}

	if response.StatusCode() != http.StatusOK {
		return fmt.Errorf("Status: %d", response.StatusCode())
	}

	return nil
}
