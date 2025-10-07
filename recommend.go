package lightnovel

type RecommendItem struct {
	GroupId    GroupID     `json:"gid"`
	Type       uint        `json:"type"`
	Title      string      `json:"title"`
	Rows       uint        `json:"rows"`
	Columns    uint        `json:"columns"`
	More       string      `json:"more"`
	MoreType   uint        `json:"more_type"`
	MoreParams any         `json:"more_params"` // can be null string "" or uint
	Items      []BookItems `json:"items"`
}

type BookItems struct {
	Id              uint          `json:"id"`
	Type            uint          `json:"type"`
	Title           string        `json:"title"`
	ActionType      uint          `json:"action_type"`
	ActionParams    uint          `json:"action_params"`
	PictureUrl      string        `json:"pic_url"`
	GroupId         GroupID       `json:"gid"`
	GroupName       string        `json:"group_name"`
	ParentGroupId   ParentGroupID `json:"parent_gid"`
	ParentGroupName string        `json:"parent_group_name"`
	Comments        uint          `json:"comments"`
	Hits            uint          `json:"hits"`
}

type RecommendRequest struct {
	UserSecurityKey
	ClassID uint `json:"class"`
}

// GetRecommendList retrieves a list of recommended items
//
// https://api.lightnovel.us/api/recom/get-recommends
func (c *Client) GetRecommendList(classID uint) ([]RecommendItem, error) {
	var req RecommendRequest
	req.UserSecurityKey = c.credentials.UserSecurityKey
	req.ClassID = classID

	resp, err := doRequest[[]RecommendItem](c, "/api/recom/get-recommends", req)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type GetRecommendRankRequest struct {
	UserSecurityKey
	ParentGroupId ParentGroupID `json:"parent_gid"`
	GroupId       GroupID       `json:"gid"`
}
type ArticleRankInfo struct {
	Rank      uint     `json:"rank"`
	ArticleId uint     `json:"aid"`
	Title     string   `json:"title"`
	Cover     string   `json:"cover"` // URL
	Comments  uint     `json:"comments"`
	Hits      uint     `json:"hits"`
	CoverType uint     `json:"cover_type"`
	Time      DateTime `json:"time"`
	SeriesId  uint     `json:"sid"`
	Banner    string   `json:"banner"` // URL
}

// GetRecommendRank retrieves article rankings
//
// https://api.lightnovel.fun/api/recom/get-ranks
func (c *Client) GetRecommendRank(parentGropuId ParentGroupID, groupId GroupID) ([]ArticleRankInfo, error) {
	req := GetRecommendRankRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		ParentGroupId:   parentGropuId,
		GroupId:         groupId,
	}

	resp, err := doRequest[[]ArticleRankInfo](c, "/api/recom/get-ranks", req)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
