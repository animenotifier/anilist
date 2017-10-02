package anilist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStreamAnime(t *testing.T) {
	APIKeyID = "akyoto-nrihb"
	APIKeySecret = "fTx1y7CwCVyQxlK54m8a8fbsEu44"

	err := Authorize()

	assert.NoError(t, err)
	assert.NotEmpty(t, AccessToken)

	allAnime := StreamAnime()

	assert.NotNil(t, allAnime)
	count := 0

	for anime := range allAnime {
		assert.NotZero(t, anime.ID)
		assert.NotEmpty(t, anime.TitleRomaji)
		assert.NotEmpty(t, anime.Type)
		assert.NotEmpty(t, anime.Link())

		count++

		if count >= 80 {
			break
		}
	}
}
