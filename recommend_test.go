package lightnovel

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRecommendList(t *testing.T) {
	client := NewClient()
	data, err := client.GetRecommendList(0)
	require.NoError(t, err)
	require.NotNil(t, data)
	require.Greater(t, len(data), 0)
	t.Logf("%+v", data)
}

func TestGetRecommendRank(t *testing.T) {
	client := NewClient()
	data, err := client.GetRecommendRank(3, 106)
	require.NoError(t, err)
	require.NotNil(t, data)
	require.Greater(t, len(data), 0)
	t.Logf("%+v", data)
}
