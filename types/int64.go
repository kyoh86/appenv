package types

import (
	"strconv"
)

type Int64Value struct {
	value int64
}

func (o *Int64Value) Value() interface{} {
	return o.value
}

func (o *Int64Value) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatInt(int64(o.value), 10)), nil
}

func (o *Int64Value) UnmarshalText(text []byte) error {
	v, err := strconv.ParseInt(string(text), 10, 64)
	if err != nil {
		return err
	}
	o.value = int64(v)
	return nil
}

func (o *Int64Value) Default() interface{} {
	return int64(0)
}

var _ Value = (*Int64Value)(nil)
