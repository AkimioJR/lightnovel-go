package lightnovel

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDiscussTopic(t *testing.T) {
	client := NewClient()
	data, err := client.GetDiscussTopic(1139088, 20, 1)
	require.NoError(t, err)
	require.NotNil(t, data)
	require.Greater(t, len(data.List), 0)
	t.Logf("%+v", data)
}
