package lightnovel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchTags(t *testing.T) {
	client := NewClient()
	data, err := client.SearchTags()
	assert.NoError(t, err)
	assert.Greater(t, len(data), 0)
	t.Logf("%+v", data)
}

func TestSearchGeneral(t *testing.T) {
	client := NewClient()
	_, err := client.SearchGeneral("刀剑", 1)
	assert.ErrorAs(t, err, &ErrNotSignedIn)
	// assert.NoError(t, err)
	// assert.NotNil(t, data)
	// assert.Greater(t, len(data.Collections), 0)
	// assert.Greater(t, len(data.Users), 0)
	// t.Logf("%+v", data)
}

func TestSearchUsers(t *testing.T) {
	client := NewClient()
	_, err := client.SearchUsers("刀剑", 1)
	assert.ErrorAs(t, err, &ErrNotSignedIn)
	// assert.NotNil(t, data)
	// assert.Greater(t, len(data.List), 0)
	// t.Logf("%+v", data)
}

func TestSearchSeries(t *testing.T) {
	client := NewClient()
	_, err := client.SearchSeries("刀剑", 1)
	assert.ErrorAs(t, err, &ErrNotSignedIn)
	// assert.NoError(t, err)
	// assert.NotNil(t, data)
	// assert.Greater(t, len(data.List), 0)
	// t.Logf("%+v", data)
}

func TestSearchLightNovels(t *testing.T) {
	client := NewClient()
	_, err := client.SearchLightNovels("刀剑", 1)
	assert.ErrorAs(t, err, &ErrNotSignedIn)
	// assert.NoError(t, err)
	// assert.NotNil(t, data)
	// assert.Greater(t, len(data.List), 0)
	// t.Logf("%+v", data)
}

func TestSearchManga(t *testing.T) {
	client := NewClient()
	_, err := client.SearchManga("刀剑", 1)
	assert.ErrorAs(t, err, &ErrNotSignedIn)
	// assert.NoError(t, err)
	// assert.NotNil(t, data)
	// assert.Greater(t, len(data.List), 0)
	// t.Logf("%+v", data)
}

func TestSearchAnime(t *testing.T) {
	client := NewClient()
	_, err := client.SearchAnime("刀剑", 1)
	assert.ErrorAs(t, err, &ErrNotSignedIn)
	// assert.NoError(t, err)
	// assert.NotNil(t, data)
	// assert.Greater(t, len(data.List), 0)
	// t.Logf("%+v", data)
}

func TestSearchNews(t *testing.T) {
	client := NewClient()
	_, err := client.SearchNews("刀剑", 1)
	assert.ErrorAs(t, err, &ErrNotSignedIn)
	// assert.NoError(t, err)
	// assert.NotNil(t, data)
	// assert.Greater(t, len(data.List), 0)
	// t.Logf("%+v", data)
}
