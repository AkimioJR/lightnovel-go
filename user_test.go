package lightnovel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var client = NewClient()

func TestLogin(t *testing.T) {
	_, err := client.Login("your_username", "your_password")
	assert.ErrorAs(t, err, &ErrLoginFailed)
}

func TestGetUserInfo(t *testing.T) {
	_, err := client.GetUserInfo()
	assert.ErrorAs(t, err, &ErrNotSignedIn)
}
