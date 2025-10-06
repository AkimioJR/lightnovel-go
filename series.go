package lightnovel

type GetSeriesInfoRequest struct {
	UserSecurityKey
	SeriesId uint `json:"sid"`
}

type SeriesInfo struct {
	Sid           uint                `json:"sid"`
	Name          string              `json:"name"`
	GroupId       uint                `json:"gid"`
	ParentGroupId uint                `json:"parent_gid"`
	Author        string              `json:"author"`
	Intro         string              `json:"intro"`
	Banner        string              `json:"banner"` // URL
	Rate          uint                `json:"rate"`
	Cover         string              `json:"cover"` // URL
	CoverType     uint                `json:"cover_type"`
	Rates         uint                `json:"rates"`
	LastTime      DateTime            `json:"last_time"`
	Hits          uint                `json:"hits"`
	Likes         uint                `json:"likes"`
	Editors       []UserProfileBase   `json:"editors"`
	Score         uint                `json:"score"`
	Characters    []any               `json:"characters"` // unknown structure
	Articles      []SeriesArticleInfo `json:"articles"`
	AlreadyFav    uint                `json:"already_fav"`
	AlreadyRate   uint                `json:"already_rate"`
	AlreadyLike   uint                `json:"already_like"`
	UserRead      UserReadInfo        `json:"user_read"`
}

type SeriesArticleInfo struct {
	Order     uint     `json:"order"`
	ArticleId uint     `json:"aid"`
	Title     string   `json:"title"`
	Banner    string   `json:"banner"` // URL
	Cover     string   `json:"cover"`  // URL
	Hits      uint     `json:"hits"`
	Comments  uint     `json:"comments"`
	CoverType uint     `json:"cover_type"`
	Time      DateTime `json:"time"`
	LastTime  DateTime `json:"last_time"`
}

type UserReadInfo struct {
	LastArticleId uint `json:"last_aid"`
}

// https://api.lightnovel.fun/api/series/get-info
func (c *Client) GetSeriesInfo(seriesId uint) (*SeriesInfo, error) {
	if c.credentials == nil {
		return nil, ErrNotSignedIn
	}
	req := GetSeriesInfoRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		SeriesId:        seriesId,
	}
	var data SeriesInfo
	err := c.doRequest("POST", "/api/series/get-info", req, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
