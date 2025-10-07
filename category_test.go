package lightnovel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCategories(t *testing.T) {
	client := NewClient()
	data, err := client.GetCategories(3)
	assert.NoError(t, err)
	assert.Greater(t, len(data), 0)
	t.Logf("%+v", data)
}

func TestGetArticleCategories(t *testing.T) {
	client := NewClient()
	data, err := client.GetArticleCategories(true, 2)
	assert.NoError(t, err)
	assert.Greater(t, len(data), 0)
	t.Logf("%+v", data)
}

func TestGetArticleByCategory(t *testing.T) {
	client := NewClient()
	data, err := client.GetArticleByCategory(3, 106, 40, 1)
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Greater(t, len(data.List), 0)
	t.Logf("%+v", data)
}
