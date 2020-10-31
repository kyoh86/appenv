package types

import (
	"strconv"
)

type Int32Value struct {
	value int32
}

func (o *Int32Value) Value() interface{} {
	return o.value
}

func (o *Int32Value) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatInt(int64(o.value), 10)), nil
}

func (o *Int32Value) UnmarshalText(text []byte) error {
	v, err := strconv.ParseInt(string(text), 10, 32)
	if err != nil {
		return err
	}
	o.value = int32(v)
	return nil
}

func (o *Int32Value) Default() interface{} {
	return int32(0)
}

var _ Value = (*Int32Value)(nil)
