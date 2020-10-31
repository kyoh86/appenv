package types

import (
	"strconv"
)

type Int8Value struct {
	value int8
}

func (o *Int8Value) Value() interface{} {
	return o.value
}

func (o *Int8Value) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatInt(int64(o.value), 10)), nil
}

func (o *Int8Value) UnmarshalText(text []byte) error {
	v, err := strconv.ParseInt(string(text), 10, 8)
	if err != nil {
		return err
	}
	o.value = int8(v)
	return nil
}

func (o *Int8Value) Default() interface{} {
	return int8(0)
}

var _ Value = (*Int8Value)(nil)
