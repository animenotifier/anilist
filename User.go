package anilist

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

// User represents an AniList user.
type User struct {
	ID              int           `json:"id"`
	DisplayName     string        `json:"display_name"`
	AnimeTime       int           `json:"anime_time"`
	MangaChap       int           `json:"manga_chap"`
	About           string        `json:"about"`
	ListOrder       int           `json:"list_order"`
	AdultContent    bool          `json:"adult_content"`
	ForumHomepage   int           `json:"forum_homepage"`
	LegacyLists     bool          `json:"legacy_lists"`
	Donator         int           `json:"donator"`
	Following       bool          `json:"following"`
	ImageURLLarge   string        `json:"image_url_lge"`
	ImageURLMedium  string        `json:"image_url_med"`
	ImageURLBanner  string        `json:"image_url_banner"`
	TitleLanguage   string        `json:"title_language"`
	ScoreType       int           `json:"score_type"`
	CustomListAnime []interface{} `json:"custom_list_anime"`
	CustomListManga []interface{} `json:"custom_list_manga"`
	Stats           struct {
		StatusDistribution struct {
			Anime struct {
				Watching    int `json:"watching"`
				PlanToWatch int `json:"plan to watch"`
				Completed   int `json:"completed"`
				Dropped     int `json:"dropped"`
				OnHold      int `json:"on-hold"`
			} `json:"anime"`
			Manga struct {
				Reading    int `json:"reading"`
				PlanToRead int `json:"plan to read"`
				Completed  int `json:"completed"`
				Dropped    int `json:"dropped"`
				OnHold     int `json:"on-hold"`
			} `json:"manga"`
		} `json:"status_distribution"`
		ScoreDistribution struct {
			Anime struct {
				Num10  int `json:"10"`
				Num20  int `json:"20"`
				Num30  int `json:"30"`
				Num40  int `json:"40"`
				Num50  int `json:"50"`
				Num60  int `json:"60"`
				Num70  int `json:"70"`
				Num80  int `json:"80"`
				Num90  int `json:"90"`
				Num100 int `json:"100"`
			} `json:"anime"`
			Manga struct {
				Num10  int `json:"10"`
				Num20  int `json:"20"`
				Num30  int `json:"30"`
				Num40  int `json:"40"`
				Num50  int `json:"50"`
				Num60  int `json:"60"`
				Num70  int `json:"70"`
				Num80  int `json:"80"`
				Num90  int `json:"90"`
				Num100 int `json:"100"`
			} `json:"manga"`
		} `json:"score_distribution"`
		FavouriteGenres struct {
			Action       int `json:"Action"`
			Comedy       int `json:"Comedy"`
			Supernatural int `json:"Supernatural"`
			Fantasy      int `json:"Fantasy"`
			Romance      int `json:"Romance"`
			Drama        int `json:"Drama"`
			SciFi        int `json:"Sci-Fi"`
			Adventure    int `json:"Adventure"`
			Mystery      int `json:"Mystery"`
			SliceOfLife  int `json:"Slice of Life"`
		} `json:"favourite_genres"`
		ActivityHistory struct {
			Two0160831 struct {
				Count int    `json:"count"`
				Color string `json:"color"`
			} `json:"2016-08-31"`
		} `json:"activity_history"`
		ActivityHistoryTotal int `json:"activity_history_total"`
		ListScores           struct {
			Anime struct {
				Mean              int `json:"mean"`
				StandardDeviation int `json:"standard_deviation"`
			} `json:"anime"`
			Manga struct {
				Mean              int `json:"mean"`
				StandardDeviation int `json:"standard_deviation"`
			} `json:"manga"`
		} `json:"list_scores"`
	} `json:"stats"`
	AdvancedRating      bool          `json:"advanced_rating"`
	AdvancedRatingNames []interface{} `json:"advanced_rating_names"`
	Notifications       int           `json:"notifications"`
	AiringNotifications int           `json:"airing_notifications"`
	UpdatedAt           int           `json:"updated_at"`
}

// GetUser ...
func GetUser(userName string) (*User, error) {
	if userName == "" {
		return nil, errors.New("Anilist username is empty")
	}

	request := gorequest.New().Get("https://anilist.co/api/user/" + userName + "?access_token=" + AccessToken)

	user := &User{}
	resp, _, errs := request.EndStruct(user)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Invalid status code: %d", resp.StatusCode)
	}

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return user, nil
}
