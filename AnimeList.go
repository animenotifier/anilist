package anilist

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

// AnimeList ...
type AnimeList struct {
	ID          int    `json:"id"`
	DisplayName string `json:"display_name"`
	Lists       struct {
		Completed   []*AnimeListItem `json:"completed"`
		Watching    []*AnimeListItem `json:"watching"`
		PlanToWatch []*AnimeListItem `json:"plan_to_watch"`
		Dropped     []*AnimeListItem `json:"dropped"`
		OnHold      []*AnimeListItem `json:"on_hold"`
	} `json:"lists"`
	CustomLists interface{} `json:"custom_lists"`
}

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
