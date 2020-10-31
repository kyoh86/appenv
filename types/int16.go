package types

import (
	"strconv"
)

type Int16Value struct {
	value int16
}

func (o *Int16Value) Value() interface{} {
	return o.value
}

func (o *Int16Value) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatInt(int64(o.value), 10)), nil
}

func (o *Int16Value) UnmarshalText(text []byte) error {
	v, err := strconv.ParseInt(string(text), 10, 16)
	if err != nil {
		return err
	}
	o.value = int16(v)
	return nil
}

func (o *Int16Value) Default() interface{} {
	return int16(0)
}

var _ Value = (*Int16Value)(nil)
