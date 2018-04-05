package anilist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByName(t *testing.T) {
	user, err := GetUserByName("Akyoto")
	assert.NoError(t, err)
	assert.Equal(t, 12647, user.ID)
}
