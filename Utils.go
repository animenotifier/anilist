package anilist

// GetAniListIDByMALID ...
func GetAniListIDByMALID(malID int) (int, error) {
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
			MALID: malID,
			Type:  "ANIME",
		},
	}

	// Query response
	response := new(struct {
		Data struct {
			Media struct {
				ID int `json:"id"`
			} `json:"Media"`
		} `json:"data"`
	})

	err := Query(body, response)

	if err != nil {
		return 0, err
	}

	return response.Data.Media.ID, nil
}
