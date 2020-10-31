package types

import (
	"strconv"
)

type Uint64Value struct {
	value uint64
}

func (o *Uint64Value) Value() interface{} {
	return o.value
}

func (o *Uint64Value) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatUint(o.value, 10)), nil
}

func (o *Uint64Value) UnmarshalText(text []byte) error {
	v, err := strconv.ParseUint(string(text), 10, 64)
	if err != nil {
		return err
	}
	o.value = uint64(v)
	return nil
}

func (o *Uint64Value) Default() interface{} {
	return uint64(0)
}

var _ Value = (*Uint64Value)(nil)
