package types

import (
	"strconv"
)

type Uint32OptionBase struct {
	value uint32
}

func (o *Uint32OptionBase) Value() interface{} {
	return o.value
}

func (o *Uint32OptionBase) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatUint(uint64(o.value), 10)), nil
}

func (o *Uint32OptionBase) UnmarshalText(text []byte) error {
	v, err := strconv.ParseUint(string(text), 10, 32)
	if err != nil {
		return err
	}
	o.value = uint32(v)
	return nil
}

func (o *Uint32OptionBase) Default() interface{} {
	return uint32(0)
}

var _ Value = (*Uint32OptionBase)(nil)
