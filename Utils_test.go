package anilist_test

import (
	"testing"

	"github.com/animenotifier/anilist"
	"github.com/stretchr/testify/assert"
)

func TestGetAniListIDByMALID(t *testing.T) {
	testIDs := map[int]int{
		1:     1,     // Cowboy Bebop
		356:   356,   // Fate Stay Night
		35062: 98436, // Mahou Tsukai no Yome
	}

	for malID, anilistID := range testIDs {
		id, err := anilist.GetAniListIDByMALID(malID)
		assert.NoError(t, err)
		assert.Equal(t, anilistID, id)
	}
}
