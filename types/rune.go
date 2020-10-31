package types

import "errors"

type RuneValue struct {
	value rune
}

func (o *RuneValue) Value() interface{} {
	return o.value
}

func (o *RuneValue) MarshalText() (text []byte, err error) {
	return []byte(string([]rune{o.value})), nil
}

func (o *RuneValue) UnmarshalText(text []byte) error {
	runes := []rune(string(text))
	if len(runes) != 1 {
		return errors.New("invalid rune")
	}
	o.value = runes[0]
	return nil
}

func (o *RuneValue) Default() interface{} {
	return rune(0)
}

var _ Value = (*RuneValue)(nil)
