package anilist_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/animenotifier/anilist"
)

func TestGetAnimeList(t *testing.T) {
	user, err := anilist.GetUser("Akyoto")
	assert.Nil(t, err)
	assert.NotNil(t, user)

	animeList, err := anilist.GetAnimeList(user.ID)
	assert.Nil(t, err)
	assert.NotNil(t, animeList)

	for _, list := range animeList.Lists {
		assert.NotEqual(t, list.Name, "")

		for _, item := range list.Entries {
			assert.NotEqual(t, item.Anime.ID, 0)
			assert.NotEqual(t, item.Anime.Title.Romaji, "")
			assert.True(t, item.Progress >= 0)
			assert.True(t, item.ScoreRaw >= 0)
			assert.True(t, item.ScoreRaw <= 100)
			assert.Contains(t, []string{
				"CURRENT",
				"PLANNING",
				"COMPLETED",
				"DROPPED",
				"PAUSED",
				"REPEATING",
			}, item.Status)
		}
	}
}
