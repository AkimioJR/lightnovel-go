package lightnovel

import "net/http"

// https://api.lightnovel.fun/api/smiley/get-ver
func (c *Client) GetVersion() (uint64, error) {
	var verion uint64
	err := c.doRequest(http.MethodPost, "/api/smiley/get-ver", nil, &verion)
	if err != nil {
		return 0, err
	}
	return verion, nil
}
