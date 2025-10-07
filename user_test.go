package lightnovel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	client := NewClient()
	_, err := client.Login("your_username", "your_password")
	assert.ErrorAs(t, err, &ErrLoginFailed)
}

func TestGetUserInfo(t *testing.T) {
	client := NewClient()
	_, err := client.GetUserInfo()
	assert.ErrorAs(t, err, &ErrNotSignedIn)
}
