package anilist

import (
	"fmt"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

// GetAnimeList ...
func GetAnimeList(userName string) (*AnimeList, error) {
	request := gorequest.New().Get("https://anilist.co/api/user/" + userName + "/animelist?access_token=" + AccessToken)

	animeList := &AnimeList{}
	resp, _, errs := request.EndStruct(animeList)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Invalid status code: %d", resp.StatusCode)
	}

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return animeList, nil
}
