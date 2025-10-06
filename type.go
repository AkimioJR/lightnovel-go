package lightnovel

import "fmt"

type UintBool bool

func (b *UintBool) UnmarshalJSON(data []byte) error {
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

func (b UintBool) MarshalJSON() ([]byte, error) {
	if b {
		return []byte("1"), nil
	} else {
		return []byte("0"), nil
	}
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
