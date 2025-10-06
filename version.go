package lightnovel

// https://api.lightnovel.fun/api/smiley/get-ver
func (c *Client) GetVersion() (uint64, error) {
	var verion uint64
	err := c.doRequest("/api/smiley/get-ver", nil, &verion)
	if err != nil {
		return 0, err
	}
	return verion, nil
}
