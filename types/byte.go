package types

import (
	"strconv"
)

type ByteOptionBase struct {
	value byte
}

func (o *ByteOptionBase) Value() interface{} {
	return o.value
}

func (o *ByteOptionBase) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatInt(int64(o.value), 10)), nil
}

func (o *ByteOptionBase) UnmarshalText(text []byte) error {
	v, err := strconv.ParseInt(string(text), 10, 8)
	if err != nil {
		return err
	}
	o.value = byte(v)
	return nil
}

func (o *ByteOptionBase) Default() interface{} {
	return byte(0)
}

var _ Value = (*ByteOptionBase)(nil)
