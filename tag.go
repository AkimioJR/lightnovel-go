package lightnovel

import "fmt"

type GetArticleTagsRequest struct {
	UserSecurityKey
	ArticleId uint `json:"article_id"`
}

func (*GetArticleTagsRequest) Path() string {
	return "/api/tag/get-article-tags"
}

func (r *GetArticleTagsRequest) CacheKey() string {
	return fmt.Sprintf("tag-get-article-tags-%d", r.ArticleId)
}

type ArticleTag struct {
	Id          uint   `json:"id"`
	Word        string `json:"word"`
	ContentType string `json:"content_type"`
	Weight      uint   `json:"weight"`
	IsClickable bool   `json:"is_clickable"`
}

// GetArticleTags retrieves tags associated with a specific article.
//
// https://api.lightnovel.fun/api/tag/get-article-tags
func (c *Client) GetArticleTags(articleId uint) ([]ArticleTag, error) {
	req := GetArticleTagsRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		ArticleId:       articleId,
	}

	resp, err := doRequest[[]ArticleTag](c, &req)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
