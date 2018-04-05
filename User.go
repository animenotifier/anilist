package anilist

// User represents an AniList user.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetUser ...
func GetUser(userName string) (*User, error) {
	type Variables struct {
		Name string `json:"name"`
	}

	body := struct {
		Query     string    `json:"query"`
		Variables Variables `json:"variables"`
	}{
		Query: `
			query ($name: String) {
				User (name: $name) {
					id
					name
				}
			}
		`,
		Variables: Variables{
			Name: userName,
		},
	}

	// Query response
	response := new(struct {
		Data struct {
			User *User `json:"User"`
		} `json:"data"`
	})

	err := Query(body, response)

	if err != nil {
		return nil, err
	}

	return response.Data.User, nil
}
