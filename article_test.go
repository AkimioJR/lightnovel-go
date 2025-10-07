package lightnovel

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetArticleDetail(t *testing.T) {
	client := NewClient()

	data, err := client.GetArticleDetail(1140926, false)
	require.NoError(t, err)
	require.NotNil(t, data)
	require.NotNil(t, data.Content)

	data, err = client.GetArticleDetail(1140926, true)
	require.NoError(t, err)
	require.NotNil(t, data)
	require.Nil(t, data.Content)
}
