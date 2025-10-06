package lightnovel

import "net/http"

type GetDiscussTopicRequest struct {
	UserSecurityKey
	ArticleId uint `json:"aid"`
	PageSize  uint `json:"page_size"`
	Page      uint `json:"page"`
}

type BaseComment struct {
	TopicId   uint            `json:"tid"`
	ArticleId uint            `json:"pid"`
	UserId    uint            `json:"uid"`
	Time      DateTime        `json:"time"`
	Content   string          `json:"content"`
	Likes     uint            `json:"likes"`
	UserInfo  UserProfileBase `json:"user_info"`
}

type TopicComment struct {
	BaseComment
	LastTime  DateTime       `json:"last_time"`
	Replies   uint           `json:"replies"`
	ReplyList []ReplyComment `json:"reply_list"`
}

type ReplyComment struct {
	BaseComment
	ReplyId          uint `json:"rid"`   // 表示当前回复的ID
	ReferenceReplyId uint `json:"r_rid"` // 表示当前回复所引用的回复ID，0表示回复楼主
	ReferenceUserID  uint `json:"r_uid"` // 当前回复所引用的回复的用户ID
}

type GetDiscussTopicResponse struct {
	List []TopicComment `json:"list"`
	Host []any          `json:"host"` // unknown structure
	Page PageInfo       `json:"page_info"`
}

// https://api.lightnovel.fun/api/discuss/get-topic
func (c *Client) GetDiscussTopic(articleId uint, pageSize uint, page uint) (*GetDiscussTopicResponse, error) {
	if c.credentials == nil {
		return nil, ErrNotSignedIn
	}
	req := GetDiscussTopicRequest{
		UserSecurityKey: c.credentials.UserSecurityKey,
		ArticleId:       articleId,
		PageSize:        pageSize,
		Page:            page,
	}
	var data GetDiscussTopicResponse
	err := c.doRequest(http.MethodPost, "/api/discuss/get-topic", req, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
