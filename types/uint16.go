package types

import (
	"strconv"
)

type Uint16Value struct {
	value uint16
}

func (o *Uint16Value) Value() interface{} {
	return o.value
}

func (o *Uint16Value) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatUint(uint64(o.value), 10)), nil
}

func (o *Uint16Value) UnmarshalText(text []byte) error {
	v, err := strconv.ParseUint(string(text), 10, 16)
	if err != nil {
		return err
	}
	o.value = uint16(v)
	return nil
}

func (o *Uint16Value) Default() interface{} {
	return uint16(0)
}

var _ Value = (*Uint16Value)(nil)
