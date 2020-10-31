package types

import "errors"

type RuneOptionBase struct {
	value rune
}

func (o *RuneOptionBase) Value() interface{} {
	return o.value
}

func (o *RuneOptionBase) MarshalText() (text []byte, err error) {
	return []byte(string([]rune{o.value})), nil
}

func (o *RuneOptionBase) UnmarshalText(text []byte) error {
	runes := []rune(string(text))
	if len(runes) != 1 {
		return errors.New("invalid rune")
	}
	o.value = runes[0]
	return nil
}

func (o *RuneOptionBase) Default() interface{} {
	return rune(0)
}

var _ Value = (*RuneOptionBase)(nil)
