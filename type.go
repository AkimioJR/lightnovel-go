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
