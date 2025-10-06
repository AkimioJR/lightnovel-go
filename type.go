package lightnovel

import (
	"encoding/json"
	"fmt"
	"time"
)

type Bool bool

func (b *Bool) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "0":
		*b = false
	case "1":
		*b = true
	default:
		return fmt.Errorf("invalid boolean value: %s", string(data))
	}
	return nil
}

func (b Bool) MarshalJSON() ([]byte, error) {
	if b {
		return []byte("1"), nil
	} else {
		return []byte("0"), nil
	}
}

type DateTime time.Time // "2006-01-02 15:04:05" time.DateTime format

func (dt *DateTime) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	if s == "0000-00-00 00:00:00" {
		*dt = DateTime(time.Time{})
		return nil
	} else {
		t, err := time.Parse(time.DateTime, s)
		if err != nil {
			return err
		}
		*dt = DateTime(t)
		return nil
	}
}

func (dt DateTime) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(dt).Format(`"2006-01-02 15:04:05"`)), nil
}

const UnknownStr = "unknown"

type PlatformType uint8

const (
	PlatformPC PlatformType = iota
	PlatformIOS
	PlatformAndroid
)

func (p PlatformType) String() string {
	switch p {
	case PlatformPC:
		return "pc"
	case PlatformIOS:
		return "ios"
	case PlatformAndroid:
		return "android"
	default:
		return UnknownStr
	}
}

func (p *PlatformType) MarshalJSON() ([]byte, error) {
	s := p.String()
	if s == UnknownStr {
		return nil, fmt.Errorf("invalid platform value")
	}
	return fmt.Appendf(nil, `"%s"`, s), nil
}

type ClientType uint8

const (
	ClientApp ClientType = iota
	ClientWeb
)

func (c ClientType) String() string {
	switch c {
	case ClientApp:
		return "app"
	case ClientWeb:
		return "web"
	default:
		return UnknownStr
	}
}

func (c *ClientType) MarshalJSON() ([]byte, error) {
	s := c.String()
	if s == UnknownStr {
		return nil, fmt.Errorf("invalid client value")
	}
	return fmt.Appendf(nil, `"%s"`, s), nil
}

type GenderType uint8

const (
	GenderUnknown GenderType = iota
	GenderMale
	GenderFemale
)

func (g *GenderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(uint8(*g))
}

type ContentType uint8

const (
	ContentGeneral    ContentType = iota // 综合
	ContentUser                          // 用户
	ContentSeries                        // 集合
	ContentNews                          // 资讯
	ContentAnime                         // 动漫
	ContentManga                         // 漫画
	ContentLightNovel                    // 轻小说
)
