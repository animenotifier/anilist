package anilist

// AnimeListItem ...
type AnimeListItem struct {
	ID        int     `json:"id"`
	Status    string  `json:"status"`
	Score     float64 `json:"score"`
	ScoreRaw  int     `json:"scoreRaw"`
	Progress  int     `json:"progress"`
	Repeat    int     `json:"repeat"`
	Private   bool    `json:"private"`
	Notes     string  `json:"notes"`
	Anime     *Anime  `json:"media"`
	StartedAt struct {
		Year  interface{} `json:"year"`
		Month interface{} `json:"month"`
		Day   interface{} `json:"day"`
	} `json:"startedAt"`
	CompletedAt struct {
		Year  interface{} `json:"year"`
		Month interface{} `json:"month"`
		Day   interface{} `json:"day"`
	} `json:"completedAt"`
	UpdatedAt int `json:"updatedAt"`
	CreatedAt int `json:"createdAt"`
}
