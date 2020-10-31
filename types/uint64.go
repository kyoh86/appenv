package types

import (
	"strconv"
)

type Uint64OptionBase struct {
	value uint64
}

func (o *Uint64OptionBase) Value() interface{} {
	return o.value
}

func (o *Uint64OptionBase) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatUint(o.value, 10)), nil
}

func (o *Uint64OptionBase) UnmarshalText(text []byte) error {
	v, err := strconv.ParseUint(string(text), 10, 64)
	if err != nil {
		return err
	}
	o.value = uint64(v)
	return nil
}

func (o *Uint64OptionBase) Default() interface{} {
	return uint64(0)
}

var _ Value = (*Uint64OptionBase)(nil)
