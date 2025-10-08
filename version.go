package lightnovel

type GetVersionRequest struct{}

func (*GetVersionRequest) Path() string {
	return "/api/smiley/get-ver"
}
func (r *GetVersionRequest) CacheKey() string {
	return "smiley-get-ver"
}

// Version represents the API version information.
//
// https://api.lightnovel.fun/api/smiley/get-ver
func (c *Client) GetVersion() (uint64, error) {
	req := GetVersionRequest{}
	resp, err := doRequest[uint64](c, &req)
	if err != nil {
		return 0, err
	}
	return resp.Data, nil
}
