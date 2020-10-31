package types

import (
	"strconv"
)

type UintOptionBase struct {
	value uint
}

func (o *UintOptionBase) Value() interface{} {
	return o.value
}

func (o *UintOptionBase) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatUint(uint64(o.value), 10)), nil
}

func (o *UintOptionBase) UnmarshalText(text []byte) error {
	v, err := strconv.ParseUint(string(text), 10, strconv.IntSize)
	if err != nil {
		return err
	}
	o.value = uint(v)
	return nil
}

func (o *UintOptionBase) Default() interface{} {
	return uint(0)
}

var _ Value = (*UintOptionBase)(nil)
