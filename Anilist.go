package anilist

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

// GetAnimeList ...
func GetAnimeList(userName string) (*AnimeList, error) {
	if userName == "" {
		return nil, errors.New("Anilist username is empty")
	}

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
