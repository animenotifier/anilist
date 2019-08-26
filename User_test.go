package anilist_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/animenotifier/anilist"
)

func TestGetUser(t *testing.T) {
	user, err := anilist.GetUser("Akyoto")
	assert.Nil(t, err)
	assert.Equal(t, 12647, user.ID)
}
