package anilist

import "strconv"

// Anime ...
type Anime struct {
	ID    int `json:"id"`
	MALID int `json:"idMal"`
	Title struct {
		Romaji  string `json:"romaji"`
		English string `json:"english"`
		Native  string `json:"native"`
	} `json:"title"`
	EpisodeCount int `json:"episodes"`
}

// Link returns the permalink to that anime.
func (anime *Anime) Link() string {
	return "https://anilist.co/anime/" + strconv.Itoa(anime.ID)
}
