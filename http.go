package lightnovel

import (
	"bytes"
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

type BaseResponse struct {
	Code      uint   `json:"code"`
	TimeStamp uint64 `json:"t"`
}
type DataResponse[T any] struct {
	Data T `json:"data"`
}

type Response[T any] struct {
	BaseResponse
	DataResponse[T]
}

func (r *Response[T]) UnmarshalJSON(data []byte) error {
	var base BaseResponse
	err := json.Unmarshal(data, &base)
	if err != nil {
		return err
	}
	r.BaseResponse = base

	if base.Code != 0 { // No need to unmarshal Data if Code is not 0 (success)
		return nil
	}
	var dr DataResponse[T]
	err = json.Unmarshal(data, &dr)
	if err != nil {
		return err
	}
	r.DataResponse = dr
	return nil
}

func doRequest[T any](c *Client, path string, data any) (*Response[T], error) {
	reqData := c.newRequest(data)
	reqBody, err := reqData.Json()
	if err != nil {
		return nil, fmt.Errorf("create request body failed: %w", err)
	}

	url := c.api + path
	req, err := http.NewRequest(http.MethodPost, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	req.Header.Set("user-agent", c.ua)
	req.Header.Set("content-type", "application/json")
	if reqData.GZ {
		req.Header.Set("accept-encoding", "gzip")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status: %s", resp.Status)
	}

	var respBody []byte
	if reqData.GZ {
		respBody, err = decompressResponse(resp)
	} else {
		respBody, err = io.ReadAll(resp.Body)
	}
	if err != nil {
		return nil, fmt.Errorf("read response body failed: %w", err)
	}

	// fmt.Println(string(b))
	var r Response[T]
	err = json.Unmarshal(respBody, &r)
	if err != nil {
		return nil, fmt.Errorf("decode response failed: %w", err)
	}
	switch r.Code {
	case 0:
		// Success
	case 5:
		return nil, ErrNotSignedIn
	default:
		return nil, fmt.Errorf("lightnovel api error: code %d", r.Code)
	}

	return &r, nil
}
