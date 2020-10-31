package types

import (
	"strconv"
)

type Uint32Value struct {
	value uint32
}

func (o *Uint32Value) Value() interface{} {
	return o.value
}

func (o *Uint32Value) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatUint(uint64(o.value), 10)), nil
}

func (o *Uint32Value) UnmarshalText(text []byte) error {
	v, err := strconv.ParseUint(string(text), 10, 32)
	if err != nil {
		return err
	}
	o.value = uint32(v)
	return nil
}

func (o *Uint32Value) Default() interface{} {
	return uint32(0)
}

var _ Value = (*Uint32Value)(nil)
