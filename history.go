package lightnovel

import "net/http"

type AddHistoryRequest struct {
	UserSecurityKey
	FId     uint `json:"fid"`
	ClassId uint `json:"class"`
}

// https://api.lightnovel.fun/api/history/add-history
func (c *Client) AddHistory(fId uint, classId uint) error {
	req := AddHistoryRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		FId:             fId,
		ClassId:         classId,
	}
	return c.doRequest(http.MethodPost, "/api/history/add-history", req, nil)
}
