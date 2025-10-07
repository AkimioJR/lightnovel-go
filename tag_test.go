package lightnovel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetArticleTags(t *testing.T) {
	client := NewClient()
	data, err := client.GetArticleTags(1140926)
	assert.NoError(t, err)
	assert.Greater(t, len(data), 0)
	t.Logf("%+v", data)
}
