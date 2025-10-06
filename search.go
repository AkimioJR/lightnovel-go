package lightnovel

import "net/http"

type Alias struct {
	ID    uint   `json:"id"`
	Alias string `json:"alias"`
}

// https://api.lightnovel.fun/api/search/get-search-tags
func (c *Client) SearchTags() ([]Alias, error) {
	var tags []Alias
	err := c.doRequest(http.MethodPost, "/api/search/get-search-tags", nil, &tags)
	if err != nil {
		return nil, err
	}
	return tags, nil
}
