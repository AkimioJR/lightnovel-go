package lightnovel

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	GZ          Bool         `json:"gz"`
	Isencrypted Bool         `json:"is_encrypted"`
	Client      ClientType   `json:"client"`
	Platform    PlatformType `json:"platform"`
	Data        any          `json:"d"`
	VersionName string       `json:"ver_name"`
	VersionCode uint         `json:"ver_code"`
	Sign        string       `json:"sign"`
}

func (c *Client) newRequest(data any) *Request {
	return &Request{
		GZ:          c.GZip,
		Isencrypted: c.Encrypted,
		Client:      c.Client,
		Platform:    c.Platform,
		VersionName: c.VersionName,
		VersionCode: c.VersionCode,
		Sign:        c.Sign,

		Data: data,
	}
}

func (r *Request) Json() (*bytes.Reader, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

type Response struct {
	Code      uint   `json:"code"`
	Data      any    `json:"data"`
	TimeStamp uint64 `json:"t"`
}

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

func (c *Client) doRequest(path string, data any, result any) error {
	var reqBody io.Reader
	if data == nil {
		reqBody = nil
	} else {
		var err error
		reqBody, err = c.newRequest(data).Json()
		if err != nil {
			return fmt.Errorf("create request body failed: %w", err)
		}
	}

	url := c.api + path
	req, err := http.NewRequest(http.MethodPost, url, reqBody)
	if err != nil {
		return fmt.Errorf("create request failed: %w", err)
	}

	req.Header.Set("user-agent", c.ua)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept-encoding", "gzip")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status: %s", resp.Status)
	}

	var decodeBody io.Reader
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		// fmt.Println("gzip encoded response")
		gr, err := gzip.NewReader(resp.Body)
		if err != nil {
			return fmt.Errorf("gzip reader error: %w", err)
		}
		defer gr.Close()
		decodeBody = gr

	default:
		decodeBody = resp.Body
	}
	b, err := io.ReadAll(decodeBody)
	if err != nil {
		return fmt.Errorf("read response body failed: %w", err)
	}

	// fmt.Println(string(b))
	var r Response
	err = json.Unmarshal(b, &r)
	if err != nil {
		return fmt.Errorf("decode response failed: %w", err)
	}
	switch r.Code {
	case 0:
		// Success
	case 5:
		return ErrNotSignedIn
	default:
		return fmt.Errorf("lightnovel api error: code %d", r.Code)
	}

	respData, err := json.Marshal(r.Data)
	if err != nil {
		return fmt.Errorf("marshal response data failed: %w", err)
	}

	if err := json.Unmarshal(respData, result); err != nil {
		return fmt.Errorf("unmarshal response data failed: %w", err)
	}

	return nil
}

func (c *Client) SetUserCredentials(uid uint, securityKey string) {
	c.credentials.UserUID.UID = uid
	c.credentials.UserSecurityKey.SecurityKey = securityKey
}
