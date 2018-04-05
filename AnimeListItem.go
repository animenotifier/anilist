package anilist

// AnimeListItem ...
type AnimeListItem struct {
	ID                    int         `json:"id"`
	Score                 float64     `json:"score"`
	ScoreRaw              int         `json:"scoreRaw"`
	Progress              int         `json:"progress"`
	ProgressVolumes       interface{} `json:"progressVolumes"`
	Repeat                int         `json:"repeat"`
	Private               bool        `json:"private"`
	Priority              int         `json:"priority"`
	Notes                 interface{} `json:"notes"`
	HiddenFromStatusLists bool        `json:"hiddenFromStatusLists"`
	StartedAt             struct {
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
	Media     struct {
		ID    int `json:"id"`
		Title struct {
			UserPreferred string `json:"userPreferred"`
		} `json:"title"`
	} `json:"media"`
}
