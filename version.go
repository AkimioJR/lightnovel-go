package lightnovel

// Version represents the API version information.
//
// https://api.lightnovel.fun/api/smiley/get-ver
func (c *Client) GetVersion() (uint64, error) {
	resp, err := doRequest[uint64](c, "/api/smiley/get-ver", nil)
	if err != nil {
		return 0, err
	}
	return resp.Data, nil
}
