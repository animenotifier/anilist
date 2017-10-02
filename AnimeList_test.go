package anilist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnimeList(t *testing.T) {
	APIKeyID = "akyoto-nrihb"
	APIKeySecret = "fTx1y7CwCVyQxlK54m8a8fbsEu44"

	err := Authorize()

	assert.NoError(t, err)
	assert.NotEmpty(t, AccessToken)

	animeList, err := GetAnimeList("Akyoto")

	assert.NoError(t, err)
	assert.NotNil(t, animeList)
	assert.NotEmpty(t, animeList.Lists.Completed)
}
