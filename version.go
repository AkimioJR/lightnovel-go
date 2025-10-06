package lightnovel

import "net/http"

// https://api.lightnovel.fun/api/smiley/get-ver
func (c *Client) GetVersion() (uint64, error) {
	var verionString uint64
	err := c.doRequest(http.MethodPost, "/api/smiley/get-ver", nil, &verionString)
	if err != nil {
		return 0, err
	}
	return verionString, nil
}
