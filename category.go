package lightnovel

import "fmt"

type GetCategoriesRequest struct {
	UserSecurityKey
	ParentGroupId ParentGroupID `json:"parent_gid"`
}

func (*GetCategoriesRequest) Path() string {
	return "/api/category/get-categories"
}

func (r *GetCategoriesRequest) CacheKey() string {
	return "category-get-categories" + fmt.Sprintf("-%d", r.ParentGroupId)
}

type CategoryInfo struct {
	GroupId    GroupID  `json:"gid"`
	Name       string   `json:"name"`
	PictureUrl string   `json:"pic"`
	LastTime   DateTime `json:"last_time"`
}

// GetCategories retrieves categories
//
// https://api.lightnovel.fun/api/category/get-categories
func (c *Client) GetCategories(parentGroupId ParentGroupID) ([]CategoryInfo, error) {
	req := GetCategoriesRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		ParentGroupId:   parentGroupId,
	}

	resp, err := doRequest[[]CategoryInfo](c, &req)
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

func (*GetArticleCategoriesRequest) Path() string {
	return "/api/category/get-article-cates"
}

func (r *GetArticleCategoriesRequest) CacheKey() string {
	return fmt.Sprintf("category-get-article-cates-%t-%d", r.Cache, r.Depth)
}

type ParentGroupCategoryInfo struct {
	GroupId   GroupID             `json:"gid"`
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

// GetArticleCategories retrieves article categories with optional caching and depth
//
// depth: 1 -> ParentGroupCategoryInfo: ParentGroupCategoryInfo.Items == nil
//
// depth: 2 -> GroupCategoryInfo: ParentGroupCategoryInfo.Items != nil
//
// https://api.lightnovel.fun/api/category/get-article-cates
func (c *Client) GetArticleCategories(cache bool, depth uint) ([]ParentGroupCategoryInfo, error) {
	req := GetArticleCategoriesRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		Cache:           cache,
		Depth:           depth,
	}

	resp, err := doRequest[[]ParentGroupCategoryInfo](c, &req)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type GetArticleByCategoryRequest struct {
	UserSecurityKey
	ParentGroupId ParentGroupID `json:"parent_gid"`
	GroupId       GroupID       `json:"gid"`
	PageSize      uint          `json:"pageSize"`
	Page          uint          `json:"page"`
}

func (*GetArticleByCategoryRequest) Path() string {
	return "/api/category/get-article-by-cate"
}

func (r *GetArticleByCategoryRequest) CacheKey() string {
	return fmt.Sprintf("category-get-article-by-cate-%d-%d-%d-%d", r.ParentGroupId, r.GroupId, r.PageSize, r.Page)
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
	GroupId   GroupID  `json:"gid"`
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

// GetArticleByCategory retrieves articles under a specific category with pagination
//
// https://api.lightnovel.fun/api/category/get-article-by-cate
func (c *Client) GetArticleByCategory(parentGroupId ParentGroupID, groupId GroupID, pageSize uint, page uint) (*GetArticleByCategoryResponse, error) {
	req := GetArticleByCategoryRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		ParentGroupId:   parentGroupId,
		GroupId:         groupId,
		PageSize:        pageSize,
		Page:            page,
	}

	resp, err := doRequest[GetArticleByCategoryResponse](c, &req)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
