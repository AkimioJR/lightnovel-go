package lightnovel

type AddHistoryRequest struct {
	UserSecurityKey
	ClassId uint `json:"class"`
	FId     uint `json:"fid"` // article id
}

func (*AddHistoryRequest) Path() string {
	return "/api/history/add-history"
}

func (r *AddHistoryRequest) CacheKey() string {
	return ""
}

// AddHistory adds a history record
//
// # Need Login
//
// https://api.lightnovel.fun/api/history/add-history
func (c *Client) AddHistory(classId uint, fId uint) error {
	req := AddHistoryRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		ClassId:         classId,
		FId:             fId,
	}
	_, err := doRequest[struct{}](c, &req)
	return err
}
