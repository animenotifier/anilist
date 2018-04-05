package anilist

import "github.com/aerogo/http/client"

var headers = client.Headers{
	"User-Agent":   "notify.moe testing",
	"Content-Type": "application/json",
}

// Query queries the AniList GraphQL API.
func Query(body interface{}) *client.Client {
	return client.Post("https://graphql.anilist.co").Headers(headers).BodyJSON(body)
}
