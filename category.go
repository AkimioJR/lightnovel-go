package lightnovel

import "net/http"

type GetArticleByCategoryRequest struct {
	UserSecurityKey
	ParentGropuId uint `json:"parent_gid"`
	GroupId       uint `json:"gid"`
	PageSize      uint `json:"pageSize"`
	Page          uint `json:"page"`
}

type CategoryInfo struct {
	ArticleId uint   `json:"aid"`
	Banner    string `json:"banner"` // URL
	Cover     string `json:"cover"`  // URL
	Title     string `json:"title"`
	Uid       uint   `json:"uid"`
	Hits      uint   `json:"hits"`
	Time      string `json:"time"`      // "2025-10-04 04:15:40" time.DateTime format
	LastTime  string `json:"last_time"` // "2025-10-04 04:15:40" time.DateTime format
	Comments  uint   `json:"comments"`
	GroupId   uint   `json:"gid"`
	GroupName string `json:"group_name"`
	CoverType uint   `json:"cover_type"`

	// 作者信息
	Author string `json:"author"`
	Avatar string `json:"avatar"`
	IsTop  uint   `json:"is_top"`

	// 仅多章节小说包含
	SeriesId   uint    `json:"sid"` // 非多章节小说默认是 0
	SeriesName *string `json:"series_name,omitempty"`
}

type PageInfo struct {
	Count        uint     `json:"count"`
	Size         uint     `json:"size"`
	Current      uint     `json:"current"`
	Previous     uint     `json:"prev"`
	Next         uint     `json:"next"`
	HasPrevious  UintBool `json:"has_prev"`
	HasNext      UintBool `json:"has_next"`
	Model        uint     `json:"model"`
	SupportModel []uint   `json:"support_model"`
}

type GetArticleByCategoryResponse struct {
	List     []CategoryInfo `json:"list"`
	PageInfo PageInfo       `json:"page_info"`
}

// https://api.lightnovel.fun/api/category/get-article-by-cate
func (c *Client) GetArticleByCategory(parentGropuId uint, groupId uint, pageSize uint, page uint) (*GetArticleByCategoryResponse, error) {
	if c.credentials == nil {
		return nil, ErrNotSignedIn
	}
	req := GetArticleByCategoryRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		ParentGropuId:   parentGropuId,
		GroupId:         groupId,
		PageSize:        pageSize,
		Page:            page,
	}
	var data GetArticleByCategoryResponse
	err := c.doRequest(http.MethodPost, "/api/category/get-article-by-cate", req, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
