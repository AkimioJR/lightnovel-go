package lightnovel

type GetArticleTagsRequest struct {
	UserSecurityKey
	ArticleId uint `json:"article_id"`
}

type ArticleTag struct {
	Id          uint   `json:"id"`
	Word        string `json:"word"`
	ContentType string `json:"content_type"`
	Weight      uint   `json:"weight"`
	IsClickable bool   `json:"is_clickable"`
}

// https://api.lightnovel.fun/api/tag/get-article-tags
func (c *Client) GetArticleTags(articleId uint) ([]ArticleTag, error) {
	if c.credentials == nil {
		return nil, ErrNotSignedIn
	}
	req := GetArticleTagsRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		ArticleId:       articleId,
	}
	var data []ArticleTag
	err := c.doRequest("POST", "/api/tag/get-article-tags", req, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
