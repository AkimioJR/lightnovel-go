package lightnovel

type GetCategoriesRequest struct {
	UserSecurityKey
	ParentGroupId uint `json:"parent_gid"`
}

type CategoryInfo struct {
	GroupId    uint     `json:"gid"`
	Name       string   `json:"name"`
	PictureUrl string   `json:"pic"`
	LastTime   DateTime `json:"last_time"`
}

// https://api.lightnovel.fun/api/category/get-categories
func (c *Client) GetCategories(parentGroupId uint) ([]CategoryInfo, error) {
	req := GetCategoriesRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		ParentGroupId:   parentGroupId,
	}

	resp, err := doRequest[[]CategoryInfo](c, "/api/category/get-categories", req)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type GetArticleCategoriesRequest struct {
	UserSecurityKey
	Cache bool `json:"cache"`
	Depth uint `json:"depth"` // 2 -> GroupCategoryInfo
}

type ParentGroupCategoryInfo struct {
	GroupId   uint                `json:"gid"`
	Name      string              `json:"name"`
	Logo      string              `json:"logo"` // URL
	CoverType uint                `json:"cover_type"`
	Order     uint                `json:"order"`
	Items     []GroupCategoryInfo `json:"items"`
}

type GroupCategoryInfo struct {
	ParentGropuId uint   `json:"gid"`
	Name          string `json:"name"`
	Logo          string `json:"logo"`
	CoverType     uint   `json:"cover_type"`
	Order         uint   `json:"order"`
}

// https://api.lightnovel.fun/api/category/get-article-cates
func (c *Client) GetArticleCategories(cache bool, depth uint) ([]ParentGroupCategoryInfo, error) {
	req := GetArticleCategoriesRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		Cache:           cache,
		Depth:           depth,
	}

	resp, err := doRequest[[]ParentGroupCategoryInfo](c, "/api/category/get-article-cates", req)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type GetArticleByCategoryRequest struct {
	UserSecurityKey
	ParentGroupId uint `json:"parent_gid"`
	GroupId       uint `json:"gid"`
	PageSize      uint `json:"pageSize"`
	Page          uint `json:"page"`
}

type ArticleInfo struct {
	ArticleId uint     `json:"aid"`
	Banner    string   `json:"banner"` // URL
	Cover     string   `json:"cover"`  // URL
	Title     string   `json:"title"`
	Uid       uint     `json:"uid"`
	Hits      uint     `json:"hits"`
	Time      DateTime `json:"time"`
	LastTime  DateTime `json:"last_time"`
	Comments  uint     `json:"comments"`
	GroupId   uint     `json:"gid"`
	GroupName string   `json:"group_name"`
	CoverType uint     `json:"cover_type"`

	// 作者信息
	Author string `json:"author"`
	Avatar string `json:"avatar"` // URL
	IsTop  uint   `json:"is_top"`

	// 仅多章节小说包含
	SeriesId   uint    `json:"sid"` // 非多章节小说默认是 0
	SeriesName *string `json:"series_name,omitempty"`
}

type PageInfo struct {
	Count        uint   `json:"count"`
	Size         uint   `json:"size"`
	Current      uint   `json:"current"`
	Previous     uint   `json:"prev"`
	Next         uint   `json:"next"`
	HasPrevious  Bool   `json:"has_prev"`
	HasNext      Bool   `json:"has_next"`
	Model        uint   `json:"model"`
	SupportModel []uint `json:"support_model"`
}

type GetArticleByCategoryResponse struct {
	List     []ArticleInfo `json:"list"`
	PageInfo PageInfo      `json:"page_info"`
}

// https://api.lightnovel.fun/api/category/get-article-by-cate
func (c *Client) GetArticleByCategory(parentGroupId uint, groupId uint, pageSize uint, page uint) (*GetArticleByCategoryResponse, error) {
	req := GetArticleByCategoryRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		ParentGroupId:   parentGroupId,
		GroupId:         groupId,
		PageSize:        pageSize,
		Page:            page,
	}

	resp, err := doRequest[GetArticleByCategoryResponse](c, "/api/category/get-article-by-cate", req)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
