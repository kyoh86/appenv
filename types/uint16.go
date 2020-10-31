package types

import (
	"strconv"
)

type Uint16OptionBase struct {
	value uint16
}

func (o *Uint16OptionBase) Value() interface{} {
	return o.value
}

func (o *Uint16OptionBase) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatUint(uint64(o.value), 10)), nil
}

func (o *Uint16OptionBase) UnmarshalText(text []byte) error {
	v, err := strconv.ParseUint(string(text), 10, 16)
	if err != nil {
		return err
	}
	o.value = uint16(v)
	return nil
}

func (o *Uint16OptionBase) Default() interface{} {
	return uint16(0)
}

var _ Value = (*Uint16OptionBase)(nil)
