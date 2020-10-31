package types

import (
	"strconv"
)

type UintValue struct {
	value uint
}

func (o *UintValue) Value() interface{} {
	return o.value
}

func (o *UintValue) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatUint(uint64(o.value), 10)), nil
}

func (o *UintValue) UnmarshalText(text []byte) error {
	v, err := strconv.ParseUint(string(text), 10, strconv.IntSize)
	if err != nil {
		return err
	}
	o.value = uint(v)
	return nil
}

func (o *UintValue) Default() interface{} {
	return uint(0)
}

var _ Value = (*UintValue)(nil)
