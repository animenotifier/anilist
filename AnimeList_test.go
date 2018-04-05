package anilist_test

import (
	"testing"

	"github.com/animenotifier/anilist"
	"github.com/stretchr/testify/assert"
)

func TestGetAnimeList(t *testing.T) {
	user, err := anilist.GetUser("Akyoto")
	assert.NoError(t, err)
	assert.NotNil(t, user)

	animeList, err := anilist.GetAnimeList(user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, animeList)

	for _, list := range animeList.Lists {
		assert.NotEmpty(t, list.Name)

		for _, item := range list.Entries {
			assert.NotZero(t, item.Anime.ID)
			assert.NotEmpty(t, item.Anime.Title.Romaji)
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
