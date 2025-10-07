package lightnovel

type Tag struct {
	Id    uint   `json:"id"`
	Alias string `json:"alias"`
}

// Get host search tags
//
// https://api.lightnovel.fun/api/search/get-search-tags
func (c *Client) SearchTags() ([]Tag, error) {
	resp, err := doRequest[[]Tag](c, "/api/search/get-search-tags", nil)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

type SearchRequest struct {
	UserSecurityKey
	Query string      `json:"q"`
	Type  ContentType `json:"type"`
	Page  uint        `json:"page"`
}

type SearchUser struct {
	UserProfileBase
	HighlightedNickname string `json:"highlighted_nickname"`
}

type SearchSeries struct {
	SeriesId        uint              `json:"sid"`
	Name            string            `json:"name"`
	Order           uint              `json:"order"`
	Author          string            `json:"author"`
	Banner          string            `json:"banner"`
	Cover           string            `json:"cover"`
	CoverType       uint              `json:"cover_type"`
	Rates           uint              `json:"rates"`
	LastTime        DateTime          `json:"last_time"`
	Hits            uint              `json:"hits"`
	Likes           uint              `json:"likes"`
	GroupId         uint              `json:"gid"`
	GroupName       string            `json:"group_name"`
	ParentGroupId   uint              `json:"parent_gid"`
	ParentGroupName string            `json:"parent_group_name"`
	Editors         []UserProfileBase `json:"editors"`
	HighlightedName string            `json:"highlighted_name"`
}

type SearchArticle struct {
	UserId           uint     `json:"uid"`
	Author           string   `json:"author"`
	Avatar           string   `json:"avatar"`
	ArticleId        uint     `json:"aid"`
	Title            string   `json:"title"`
	Banner           string   `json:"banner"`
	Cover            string   `json:"cover"`
	CoverType        uint     `json:"cover_type"`
	Hits             uint     `json:"hits"`
	Comments         uint     `json:"comments"`
	Time             DateTime `json:"time"`
	GroupId          uint     `json:"gid"`
	GroupName        string   `json:"group_name"`
	ParentGid        uint     `json:"parent_gid"`
	ParentGroupName  string   `json:"parent_group_name"`
	SeriesId         uint     `json:"sid"`
	SeriesName       string   `json:"series_name"`
	HighlightedTitle string   `json:"highlighted_title"`
}

type SearchNew struct {
	UserId uint   `json:"uid"`
	Author string `json:"author"`
	Avatar string `json:"avatar"`

	ArticleId uint     `json:"aid"`
	Title     string   `json:"title"`
	Banner    string   `json:"banner"`
	Cover     string   `json:"cover"`
	CoverType uint     `json:"cover_type"`
	Hits      uint     `json:"hits"`
	Comments  uint     `json:"comments"`
	Time      DateTime `json:"time"`

	GroupId         uint   `json:"gid"`
	GroupName       string `json:"group_name"`
	ParentGroupId   uint   `json:"parent_gid"`
	ParentGroupName string `json:"parent_group_name"`

	SeriesId         uint   `json:"sid"`
	Empty            Bool   `json:"empty"`
	HighlightedTitle string `json:"highlighted_title"`
}

type SearchGeneralResponse struct {
	Games       []any          `json:"games"`
	Collections []SearchSeries `json:"collections"`
	Users       []SearchUser   `json:"users"`
	PageInfo    PageInfo       `json:"page_info"`
}

type SearchResponse[T any] struct {
	List     []T      `json:"list"`
	PageInfo PageInfo `json:"page_info"`
}

// ContentType represents the type of content to search for
//
// https://api.lightnovel.fun/api/search/search-result
func search[T any](c *Client, query string, page uint, t ContentType) (*T, error) {
	req := SearchRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		Query:           query,
		Type:            t,
		Page:            page,
	}

	resp, err := doRequest[T](c, "/api/search/search-result", req)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// SearchGeneral performs a general search
//
// # Need Login
//
// https://api.lightnovel.fun/api/search/search-result
func (c *Client) SearchGeneral(query string, page uint) (*SearchGeneralResponse, error) {
	return search[SearchGeneralResponse](c, query, page, ContentGeneral)
}

// SearchUsers searches for users
//
// # Need Login
//
// https://api.lightnovel.fun/api/search/search-result
func (c *Client) SearchUsers(query string, page uint) (*SearchResponse[UserProfileBase], error) {
	return search[SearchResponse[UserProfileBase]](c, query, page, ContentUser)
}

// SearchSeries searches for series
//
// # Need Login
//
// https://api.lightnovel.fun/api/search/search-result
func (c *Client) SearchSeries(query string, page uint) (*SearchResponse[SearchSeries], error) {
	return search[SearchResponse[SearchSeries]](c, query, page, ContentSeries)
}

// SearchLightNovels searches for lightnovels
//
// # Need Login
//
// https://api.lightnovel.fun/api/search/search-result
func (c *Client) SearchLightNovels(query string, page uint) (*SearchResponse[SearchArticle], error) {
	return search[SearchResponse[SearchArticle]](c, query, page, ContentLightNovel)
}

// SearchManga searches for manga
//
// # Need Login
//
// https://api.lightnovel.fun/api/search/search-result
func (c *Client) SearchManga(query string, page uint) (*SearchResponse[SearchArticle], error) {
	return search[SearchResponse[SearchArticle]](c, query, page, ContentManga)
}

// SearchAnime searches for anime
//
// # Need Login
//
// https://api.lightnovel.fun/api/search/search-result
func (c *Client) SearchAnime(query string, page uint) (*SearchResponse[SearchArticle], error) {
	return search[SearchResponse[SearchArticle]](c, query, page, ContentAnime)
}

// SearchNews searches for news for lightnovels, manga, anime and buangumi
//
// # Need Login
//
// https://api.lightnovel.fun/api/search/search-result
func (c *Client) SearchNews(query string, page uint) (*SearchResponse[SearchNew], error) {
	return search[SearchResponse[SearchNew]](c, query, page, ContentNews)
}
