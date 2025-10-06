package lightnovel

import "net/http"

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

// UserCredentials represents user authentication information
type UserCredentials struct {
	UserUID
	UserSecurityKey
}

type UserProfileBase struct {
	UserUID
	NickName   string        `json:"nickname"`
	Avatar     string        `json:"avatar"` // Avatar image URL
	Passer     uint          `json:"passer"`
	Gender     uint          `json:"gender"`
	Sign       string        `json:"sign"`
	Status     uint          `json:"status"`
	Banner     string        `json:"banner"`       // Banner image URL
	BanEndDate string        `json:"ban_end_date"` // Date when ban ends "000-00-00 00:00:00" time.DateTime format
	Medals     []any         `json:"medals"`       // unknown structure
	Following  uint          `json:"following"`    // Number of users this user is following
	Comments   uint          `json:"comments"`
	Favorites  uint          `json:"favorites"`
	Articles   uint          `json:"articles"`
	Level      UserLevelInfo `json:"level"`
	Balance    UserBalance   `json:"balance"`
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

// https://api.lightnovel.fun/api/user/login
func (c *Client) Login(username, password string) (*UserLoginResponse, error) {

	var data UserLoginResponse
	err := c.doRequest(http.MethodPost,
		"/api/user/login",
		LoginRequest{
			Username: username,
			Password: password,
		},
		&data,
	)
	if err != nil {
		return nil, err
	}

	c.credentials = &UserCredentials{
		UserUID:         data.UserUID,
		UserSecurityKey: data.UserSecurityKey,
	}
	return &data, nil
}
