package types

import (
	"strconv"
)

type Uint8Value struct {
	value uint8
}

func (o *Uint8Value) Value() interface{} {
	return o.value
}

func (o *Uint8Value) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatUint(uint64(o.value), 10)), nil
}

func (o *Uint8Value) UnmarshalText(text []byte) error {
	v, err := strconv.ParseUint(string(text), 10, 8)
	if err != nil {
		return err
	}
	o.value = uint8(v)
	return nil
}

func (o *Uint8Value) Default() interface{} {
	return uint8(0)
}

var _ Value = (*Uint8Value)(nil)
