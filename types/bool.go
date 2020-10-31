package types

import (
	"strconv"
)

type BoolOptionBase struct {
	value bool
}

func (o *BoolOptionBase) Value() interface{} {
	return o.value
}

func (o *BoolOptionBase) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatBool(o.value)), nil
}

func (o *BoolOptionBase) UnmarshalText(text []byte) error {
	v, err := strconv.ParseBool(string(text))
	if err != nil {
		return err
	}
	o.value = v
	return nil
}

func (o *BoolOptionBase) Default() interface{} {
	return false
}

var _ Value = (*BoolOptionBase)(nil)
