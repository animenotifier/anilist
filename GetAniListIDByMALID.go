package anilist

import (
	"strconv"
)

// GetAniListIDByMALID ...
func GetAniListIDByMALID(malID string) (string, error) {
	malIDNumber, _ := strconv.Atoi(malID)

	type Variables struct {
		MALID int    `json:"malId"`
		Type  string `json:"type"`
	}

	body := struct {
		Query     string    `json:"query"`
		Variables Variables `json:"variables"`
	}{
		Query: `
			query ($malId: Int, $type: MediaType) {
				Media(idMal: $malId, type: $type) {
					id
				}
			}
		`,
		Variables: Variables{
			MALID: malIDNumber,
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

	err := Query(body, idResponse)

	if err != nil {
		return "", err
	}

	return strconv.Itoa(idResponse.Data.Media.ID), nil
}
