package types

import (
	"strconv"
)

type IntValue struct {
	value int
}

func (o *IntValue) Value() interface{} {
	return o.value
}

func (o *IntValue) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatInt(int64(o.value), 10)), nil
}

func (o *IntValue) UnmarshalText(text []byte) error {
	v, err := strconv.ParseInt(string(text), 10, strconv.IntSize)
	if err != nil {
		return err
	}
	o.value = int(v)
	return nil
}

func (o *IntValue) Default() interface{} {
	return int(0)
}

var _ Value = (*IntValue)(nil)
