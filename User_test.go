package anilist_test

import (
	"testing"

	"github.com/animenotifier/anilist"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	user, err := anilist.GetUser("Akyoto")
	assert.NoError(t, err)
	assert.Equal(t, 12647, user.ID)
}
