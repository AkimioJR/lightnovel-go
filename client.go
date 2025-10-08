package lightnovel

import (
	"net/http"
)

type Client struct {
	api        string
	httpClient *http.Client
	ua         string

	credentials UserCredentials

	GZip        Bool
	Encrypted   Bool
	Client      ClientType
	Platform    PlatformType
	VersionName string
	VersionCode uint
	Sign        string
}

func NewClient() *Client {
	c := Client{
		api:        "https://api.lightnovel.fun",
		httpClient: &http.Client{},
		ua:         "Dart/2.10 (dart:io)",

		GZip:        false,
		Encrypted:   false,
		Client:      ClientApp,
		Platform:    PlatformIOS,
		VersionName: "0.11.51",
		VersionCode: 191,
		Sign:        "",
	}
	return &c
}

func (c *Client) SetUserAgent(ua string) {
	c.ua = ua
}

// SetAPIEndpoint sets the API endpoint URL.
//
// Default is "https://api.lightnovel.fun" application API.
//
// "https://www.lightnovel.fun/proxy" is web API.
func (c *Client) SetAPIEndpoint(api string) {
	c.api = api
}

func (c *Client) SetUserCredentials(uid uint, securityKey string) {
	c.credentials.UserUID.UID = uid
	c.credentials.UserSecurityKey.SecurityKey = securityKey
}

func (c *Client) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}
