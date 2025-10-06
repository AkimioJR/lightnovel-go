package lightnovel

type Alias struct {
	ID    uint   `json:"id"`
	Alias string `json:"alias"`
}

// https://api.lightnovel.fun/api/search/get-search-tags
func (c *Client) SearchTags() ([]Alias, error) {
	var tags []Alias
	err := c.doRequest("/api/search/get-search-tags", nil, &tags)
	if err != nil {
		return nil, err
	}
	return tags, nil
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

// https://api.lightnovel.fun/api/search/search-result
func search[T any](c *Client, query string, page uint, t ContentType) (*T, error) {
	req := SearchRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		Query:           query,
		Type:            t,
		Page:            page,
	}

	// b, _ := json.Marshal(newRequest(req))
	// fmt.Println(string(b))
	var data T
	err := c.doRequest("/api/search/search-result", req, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) SearchGeneral(query string, page uint) (*SearchGeneralResponse, error) {
	return search[SearchGeneralResponse](c, query, page, ContentGeneral)
}

func (c *Client) SearchUsers(query string, page uint) (*SearchResponse[UserProfileBase], error) {
	return search[SearchResponse[UserProfileBase]](c, query, page, ContentUser)
}

func (c *Client) SearchSeries(query string, page uint) (*SearchResponse[SearchSeries], error) {
	return search[SearchResponse[SearchSeries]](c, query, page, ContentSeries)
}

func (c *Client) SearchLightNovels(query string, page uint) (*SearchResponse[SearchArticle], error) {
	return search[SearchResponse[SearchArticle]](c, query, page, ContentLightNovel)
}

func (c *Client) SearchManga(query string, page uint) (*SearchResponse[SearchArticle], error) {
	return search[SearchResponse[SearchArticle]](c, query, page, ContentManga)
}

// https://api.lightnovel.fun/api/search/search-result
func (c *Client) SearchAnime(query string, page uint) (*SearchResponse[SearchArticle], error) {
	return search[SearchResponse[SearchArticle]](c, query, page, ContentAnime)
}

func (c *Client) SearchNews(query string, page uint) (*SearchResponse[SearchNew], error) {
	return search[SearchResponse[SearchNew]](c, query, page, ContentNews)
}
