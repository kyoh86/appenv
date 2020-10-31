package types

import (
	"strconv"
)

type IntOptionBase struct {
	value int
}

func (o *IntOptionBase) Value() interface{} {
	return o.value
}

func (o *IntOptionBase) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatInt(int64(o.value), 10)), nil
}

func (o *IntOptionBase) UnmarshalText(text []byte) error {
	v, err := strconv.ParseInt(string(text), 10, strconv.IntSize)
	if err != nil {
		return err
	}
	o.value = int(v)
	return nil
}

func (o *IntOptionBase) Default() interface{} {
	return int(0)
}

var _ Value = (*IntOptionBase)(nil)
