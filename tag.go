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
	req := GetArticleTagsRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		ArticleId:       articleId,
	}

	resp, err := doRequest[[]ArticleTag](c, "/api/tag/get-article-tags", req)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
