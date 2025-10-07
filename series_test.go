package lightnovel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSeriesInfo(t *testing.T) {
	client := NewClient()
	data, err := client.GetSeriesInfo(5253)
	assert.NoError(t, err)
	assert.NotNil(t, data)
	t.Logf("%+v", data)
}
