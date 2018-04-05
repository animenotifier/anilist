package anilist_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/animenotifier/anilist"
	"github.com/stretchr/testify/assert"
)

func TestGetAnimeList(t *testing.T) {
	user, _ := anilist.GetUser("Akyoto")
	animeList, err := anilist.GetAnimeList(user.ID)

	assert.NoError(t, err)
	assert.NotNil(t, animeList)

	pretty, _ := json.MarshalIndent(animeList, "", "\t")
	fmt.Println(string(pretty))
}
