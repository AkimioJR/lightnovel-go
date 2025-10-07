package lightnovel

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetVersion(t *testing.T) {
	client := NewClient()
	client.GZip = false
	v1, err := client.GetVersion()
	t.Logf("version: %d", v1)
	require.NoError(t, err, "no gzip")

	client.GZip = true
	v2, err := client.GetVersion()
	t.Logf("version: %d", v2)
	require.NoError(t, err, "with gzip")

	require.Equal(t, v1, v2, "版本号应该相同")
}
