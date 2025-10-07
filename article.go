package lightnovel

type ArticleDetail struct {
	ArticleId     uint             `json:"aid"`
	UserId        uint             `json:"uid"`
	Title         string           `json:"title"`
	Summary       string           `json:"summary"`
	Content       *string          `json:"content"`
	Hits          uint             `json:"hits"`
	Likes         uint             `json:"likes"`
	Coins         uint             `json:"coins"`
	Favorites     uint             `json:"favorites"`
	Comments      uint             `json:"comments"`
	Shares        uint             `json:"shares"`
	Time          DateTime         `json:"time"`
	HasPoll       Bool             `json:"has_poll"`
	Banner        string           `json:"banner"`
	OnlyPasser    Bool             `json:"only_passer"`
	Cover         string           `json:"cover"` // URL
	LastTime      DateTime         `json:"last_time"`
	Lt            DateTime         `json:"lt"`
	GroupId       uint             `json:"gid"`
	ParentGroupId uint             `json:"parent_gid"`
	SeriesId      uint             `json:"sid"`
	Author        UserProfileBase  `json:"author"`
	OtherRecoms   []any            `json:"other_recoms"` // unknown structure
	Res           PictureResources `json:"res"`          // 图片资源
	CacheVer      uint             `json:"cache_ver"`
	OnlyApp       Bool             `json:"only_app"`
	AlreadyCoin   uint             `json:"already_coin"` // 投币数量 0 -> 未投币
	AlreadyLike   Bool             `json:"already_like"`
	AlreadyFav    Bool             `json:"already_fav"`
	AlreadyFollow Bool             `json:"already_follow"`
}

type PictureResources struct {
	Ids           []string               `json:"ids"`
	ResourcesInfo map[string]PictureInfo `json:"res_info"`
}

type PictureInfo struct {
	Resid    uint   `json:"resid"`
	Width    uint   `json:"width"`
	Height   uint   `json:"height"`
	Ext      string `json:"ext"`
	Filename string `json:"filename"`
	Url      string `json:"url"`
}

type TetArticleDetailRequest struct {
	UserSecurityKey
	ArticleId uint `json:"aid"`
	NoContent bool `json:"simple"`
}

// GetArticleDetail retrieves detailed information about a specific article
// 
// https://api.lightnovel.fun/api/article/get-detail
func (c *Client) GetArticleDetail(articleId uint, noContent bool) (*ArticleDetail, error) {
	req := TetArticleDetailRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		ArticleId:       articleId,
		NoContent:       noContent,
	}
	resp, err := doRequest[ArticleDetail](c, "/api/article/get-detail", req)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
