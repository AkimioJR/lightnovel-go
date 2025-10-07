package lightnovel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddHistory(t *testing.T) {
	client := NewClient()
	err := client.AddHistory(1, 1140926)
	assert.ErrorAs(t, err, &ErrNotSignedIn)
}
