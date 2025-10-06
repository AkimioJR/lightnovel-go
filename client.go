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
	GZ          int    `json:"gz"`
	Client      string `json:"client"`
	Platform    string `json:"platform"`
	Data        any    `json:"d"`
	VersionName string `json:"ver_name"`
	VersionCode int    `json:"ver_code"`
	Sign        string `json:"sign"`
}

func newRequest(data any) *Request {
	return &Request{
		GZ:          1,
		Client:      "app",
		Platform:    "ios",
		Data:        data,
		VersionName: "0.11.51",
		VersionCode: 191,
		Sign:        "",
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
	Code      int   `json:"code"`
	Data      any   `json:"data"`
	TimeStamp int64 `json:"t"`
}

type Client struct {
	api        string
	httpClient *http.Client
	ua         string
}

func NewClient() *Client {
	c := Client{
		api:        "https://api.lightnovel.fun",
		httpClient: &http.Client{},
		ua:         "Dart/2.10 (dart:io)",
	}
	return &c
}

func (c *Client) SetUserAgent(ua string) {
	c.ua = ua
}

func (c *Client) doRequest(method string, path string, data any, result any) error {
	var reqBody io.Reader
	if method == http.MethodGet || data == nil {
		reqBody = nil
	} else {
		var err error
		reqBody, err = newRequest(data).Json()
		if err != nil {
			return fmt.Errorf("create request body failed: %w", err)
		}
	}

	url := c.api + path
	req, err := http.NewRequest(method, url, reqBody)
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

	var respBody io.Reader
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		gr, err := gzip.NewReader(resp.Body)
		if err != nil {
			return fmt.Errorf("gzip reader error: %w", err)
		}
		defer gr.Close()
		respBody = gr

	default:
		respBody = resp.Body
	}

	var r Response
	if err := json.NewDecoder(respBody).Decode(&r); err != nil {
		return fmt.Errorf("decode response failed: %w", err)
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
