package types

import (
	"strconv"
)

type ByteValue struct {
	value byte
}

func (o *ByteValue) Value() interface{} {
	return o.value
}

func (o *ByteValue) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatInt(int64(o.value), 10)), nil
}

func (o *ByteValue) UnmarshalText(text []byte) error {
	v, err := strconv.ParseInt(string(text), 10, 8)
	if err != nil {
		return err
	}
	o.value = byte(v)
	return nil
}

func (o *ByteValue) Default() interface{} {
	return byte(0)
}

var _ Value = (*ByteValue)(nil)
