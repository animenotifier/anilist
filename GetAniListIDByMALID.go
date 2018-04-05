package anilist

import (
	"fmt"
	"net/http"
	"strconv"
)

const malIDQuery = `query ($malId: Int, $type: MediaType) {
	Media(idMal: $malId, type: $type) {
		id
	}
}`

type malIDQueryBody struct {
	Query     string              `json:"query"`
	Variables malIDQueryVariables `json:"variables"`
}

type malIDQueryVariables struct {
	MALID int    `json:"malId"`
	Type  string `json:"type"`
}

// GetAniListIDByMALID ...
func GetAniListIDByMALID(malID string) (string, error) {
	malIDInteger, _ := strconv.Atoi(malID)

	body := malIDQueryBody{
		Query: malIDQuery,
		Variables: malIDQueryVariables{
			MALID: malIDInteger,
			Type:  "ANIME",
		},
	}

	// Query response
	idResponse := new(struct {
		Data struct {
			Media struct {
				ID int `json:"id"`
			} `json:"Media"`
		} `json:"data"`
	})

	response, err := Query(body).EndStruct(idResponse)

	if err != nil {
		return "", err
	}

	if response.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("Status code: %d", response.StatusCode())
	}

	return strconv.Itoa(idResponse.Data.Media.ID), nil
}
