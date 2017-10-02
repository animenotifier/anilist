package anilist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorize(t *testing.T) {
	APIKeyID = "akyoto-nrihb"
	APIKeySecret = "fTx1y7CwCVyQxlK54m8a8fbsEu44"

	err := Authorize()

	assert.NoError(t, err)
	assert.NotEmpty(t, AccessToken)
}

func TestAuthorizeFail(t *testing.T) {
	APIKeyID = "must-fail"
	APIKeySecret = "intentionally-wrong-secret"

	err := Authorize()

	assert.Error(t, err)
	assert.Empty(t, AccessToken)
}
