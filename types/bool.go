package types

import (
	"strconv"
)

type BoolValue struct {
	value bool
}

func (o *BoolValue) Value() interface{} {
	return o.value
}

func (o *BoolValue) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatBool(o.value)), nil
}

func (o *BoolValue) UnmarshalText(text []byte) error {
	v, err := strconv.ParseBool(string(text))
	if err != nil {
		return err
	}
	o.value = v
	return nil
}

func (o *BoolValue) Default() interface{} {
	return false
}

var _ Value = (*BoolValue)(nil)
