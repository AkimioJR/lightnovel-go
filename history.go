package lightnovel

type AddHistoryRequest struct {
	UserSecurityKey
	ClassId uint `json:"class"`
	FId     uint `json:"fid"`
}

// https://api.lightnovel.fun/api/history/add-history
func (c *Client) AddHistory(classId uint, fId uint) error {
	req := AddHistoryRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		ClassId:         classId,
		FId:             fId,
	}
	_, err := doRequest[any](c, "/api/history/add-history", req)
	return err
}
