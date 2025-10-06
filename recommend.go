package lightnovel

import "net/http"

type RecommendItem struct {
	Gid        int64       `json:"gid"`
	Type       int64       `json:"type"`
	Title      string      `json:"title"`
	Rows       int64       `json:"rows"`
	Columns    int64       `json:"columns"`
	More       string      `json:"more"`
	MoreType   int64       `json:"more_type"`
	MoreParams any         `json:"more_params"` // can be null string "" or uint
	Items      []BookItems `json:"items"`
}

type BookItems struct {
	Id              int64  `json:"id"`
	Type            int64  `json:"type"`
	Title           string `json:"title"`
	ActionType      int64  `json:"action_type"`
	ActionParams    int64  `json:"action_params"`
	PictureUrl      string `json:"pic_url"`
	Gid             int64  `json:"gid"`
	GroupName       string `json:"group_name"`
	ParentGid       int64  `json:"parent_gid"`
	ParentGroupName string `json:"parent_group_name"`
	Comments        int64  `json:"comments"`
	Hits            int64  `json:"hits"`
}

type RecommendRequest struct {
	UserSecurityKey
	ClassID uint `json:"class"`
}

// https://api.lightnovel.us/api/recom/get-recommends
func (c *Client) GetRecommendList(classID uint) ([]RecommendItem, error) {
	var req RecommendRequest
	if c.credentials == nil {
		return nil, ErrNotSignedIn
	}
	req.UserSecurityKey = c.credentials.UserSecurityKey
	req.ClassID = classID

	var data []RecommendItem
	err := c.doRequest(http.MethodPost, "/api/recom/get-recommends", req, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
