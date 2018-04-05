package anilist

// AnimeList ...
type AnimeList struct {
	Lists []struct {
		Name                 string           `json:"name"`
		IsCustomList         bool             `json:"isCustomList"`
		IsSplitCompletedList bool             `json:"isSplitCompletedList"`
		Entries              []*AnimeListItem `json:"entries"`
	} `json:"lists"`
	User struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Avatar struct {
			Large string `json:"large"`
		} `json:"avatar"`
		MediaListOptions struct {
			ScoreFormat string `json:"scoreFormat"`
			RowOrder    string `json:"rowOrder"`
		} `json:"mediaListOptions"`
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
						isCustomList
						isSplitCompletedList
						entries {
							... mediaListEntry
						}
					}
					user {
						id
						name 
						avatar {
							large
						}
						mediaListOptions {
							scoreFormat
							rowOrder
						}
					}
				}
			}
			
			fragment mediaListEntry on MediaList {
				id
				score
				scoreRaw: score (format: POINT_100)
				progress
				progressVolumes
				repeat
				private
				priority
				notes
				hiddenFromStatusLists
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
					title {
						userPreferred
					}
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
