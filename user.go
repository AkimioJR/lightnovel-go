package lightnovel

import (
	"fmt"
)

// ExperienceLevel represents a level in the experience system
type ExperienceLevel struct {
	Experience uint   `json:"exp"`
	Level      uint   `json:"level"`
	Name       string `json:"name"`
}

// UserLevelInfo represents user's current level information including next level requirements
type UserLevelInfo struct {
	ExperienceLevel
	NextExp uint `json:"next_exp"`
}

// UserBalance represents user's currency and balance information
type UserBalance struct {
	Coin    uint `json:"coin"`
	Balance uint `json:"balance"`
}

type UserUID struct {
	UID uint `json:"uid"`
}

type UserSecurityKey struct {
	SecurityKey string `json:"security_key"`
}

type Medals struct {
	MedalId    uint     `json:"medal_id"`
	Name       string   `json:"name"`
	Desc       string   `json:"desc"`
	Type       uint     `json:"type"`
	Equip      uint     `json:"equip"`
	Expiration DateTime `json:"expiration"`
	Img        string   `json:"img"`
}

// UserCredentials represents user authentication information
type UserCredentials struct {
	UserUID
	UserSecurityKey
}

type UserProfileBase struct {
	UserUID
	NickName   string        `json:"nickname"`
	Avatar     string        `json:"avatar"` // Avatar image URL
	Passer     Bool          `json:"passer"`
	Gender     GenderType    `json:"gender"`
	Sign       string        `json:"sign"`
	Status     Bool          `json:"status"`
	Banner     string        `json:"banner"`       // Banner image URL
	BanEndDate DateTime      `json:"ban_end_date"` // Date when ban ends
	Medals     []Medals      `json:"medals"`
	Following  uint          `json:"following"` // Number of users this user is following
	Favorites  uint          `json:"favorites"`
	Articles   uint          `json:"articles"`
	Level      UserLevelInfo `json:"level"`

	// 仅能看到自己的信息
	Comments *uint        `json:"comments"`
	Balance  *UserBalance `json:"balance"`
}

// UserProfile represents complete user profile information
type UserProfileDetail struct {
	UserProfileBase
	AllLevel []ExperienceLevel `json:"all_level"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	UserProfileBase
	UserSecurityKey
}

var ErrLoginFailed = fmt.Errorf("login failed")

// https://api.lightnovel.fun/api/user/login
func (c *Client) Login(username, password string) (*UserLoginResponse, error) {
	resp, err := doRequest[UserLoginResponse](
		c,
		"/api/user/login",
		LoginRequest{
			Username: username,
			Password: password,
		},
	)
	if err != nil {
		return nil, err
	}

	if resp.Data.UserUID.UID == 0 || resp.Data.UserSecurityKey.SecurityKey == "" {
		return nil, ErrLoginFailed
	}
	c.SetUserCredentials(resp.Data.UserUID.UID, resp.Data.UserSecurityKey.SecurityKey)
	return &resp.Data, nil
}

var ErrNotSignedIn = fmt.Errorf("user not signed in")

// https://api.lightnovel.fun/api/user/info
func (c *Client) GetUserInfo() (*UserProfileDetail, error) {
	resp, err := doRequest[UserProfileDetail](c, "/api/user/info", c.credentials)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	return &resp.Data, nil
}
