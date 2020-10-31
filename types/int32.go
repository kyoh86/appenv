package types

import (
	"strconv"
)

type Int32OptionBase struct {
	value int32
}

func (o *Int32OptionBase) Value() interface{} {
	return o.value
}

func (o *Int32OptionBase) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatInt(int64(o.value), 10)), nil
}

func (o *Int32OptionBase) UnmarshalText(text []byte) error {
	v, err := strconv.ParseInt(string(text), 10, 32)
	if err != nil {
		return err
	}
	o.value = int32(v)
	return nil
}

func (o *Int32OptionBase) Default() interface{} {
	return int32(0)
}

var _ Value = (*Int32OptionBase)(nil)
