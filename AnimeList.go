package anilist

// AnimeList ...
type AnimeList struct {
	Lists []struct {
		Name    string           `json:"name"`
		Entries []*AnimeListItem `json:"entries"`
	} `json:"lists"`
	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
}

// GetAnimeList ...
func GetAnimeList(userID int) (*AnimeList, error) {
	type Variables struct {
		ID       int    `json:"id"`
		ListType string `json:"listType"`
	}

	body := struct {
		Query     string    `json:"query"`
		Variables Variables `json:"variables"`
	}{
		Query: `
			query ($id: Int!, $listType: MediaType) {
				MediaListCollection (userId: $id, type: $listType) {
					lists {
						name
						entries {
							... mediaListEntry
						}
					}
					user {
						id
						name
					}
				}
			}
			
			fragment mediaListEntry on MediaList {
				id
				score
				scoreRaw: score (format: POINT_100)
				progress
				repeat
				private
				notes
				status
				startedAt {
					year
					month
					day
				}
				completedAt {
					year
					month
					day
				}
				updatedAt
				createdAt
				media {
					id
					idMal
					title {
						romaji
						english
						native
					}
					episodes
				}
			}
		`,
		Variables: Variables{
			ID:       userID,
			ListType: "ANIME",
		},
	}

	// Query response
	response := new(struct {
		Data struct {
			MediaListCollection *AnimeList `json:"MediaListCollection"`
		} `json:"data"`
	})

	err := Query(body, &response)

	if err != nil {
		return nil, err
	}

	return response.Data.MediaListCollection, nil
}
