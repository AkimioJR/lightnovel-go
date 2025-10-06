package lightnovel

import "net/http"

type RecommendItem struct {
	GroupId    uint        `json:"gid"`
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
	Id              uint   `json:"id"`
	Type            uint   `json:"type"`
	Title           string `json:"title"`
	ActionType      uint   `json:"action_type"`
	ActionParams    uint   `json:"action_params"`
	PictureUrl      string `json:"pic_url"`
	GroupId         uint   `json:"gid"`
	GroupName       string `json:"group_name"`
	ParentGroupId   uint   `json:"parent_gid"`
	ParentGroupName string `json:"parent_group_name"`
	Comments        uint   `json:"comments"`
	Hits            uint   `json:"hits"`
}

type RecommendRequest struct {
	UserSecurityKey
	ClassID uint `json:"class"`
}

// https://api.lightnovel.us/api/recom/get-recommends
func (c *Client) GetRecommendList(classID uint) ([]RecommendItem, error) {
	var req RecommendRequest
	req.UserSecurityKey = c.credentials.UserSecurityKey
	req.ClassID = classID

	var data []RecommendItem
	err := c.doRequest(http.MethodPost, "/api/recom/get-recommends", req, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type GetRecommendRankRequest struct {
	UserSecurityKey
	ParentGroupId uint `json:"parent_gid"`
	GroupId       uint `json:"gid"`
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

// https://api.lightnovel.fun/api/recom/get-ranks
func (c *Client) GetRecommendRank(parentGropuId uint, groupId uint) ([]ArticleRankInfo, error) {
	req := GetRecommendRankRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		ParentGroupId:   parentGropuId,
		GroupId:         groupId,
	}
	var data []ArticleRankInfo
	err := c.doRequest(http.MethodPost, "/api/recom/get-ranks", req, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
