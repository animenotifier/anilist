package anilist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	APIKeyID = "akyoto-nrihb"
	APIKeySecret = "fTx1y7CwCVyQxlK54m8a8fbsEu44"

	err := Authorize()
	assert.NoError(t, err)
	assert.NotEmpty(t, AccessToken)

	user, err := GetUser("Akyoto")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ImageURLLarge)
}
