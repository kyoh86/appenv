package types

import (
	"strconv"
)

type Float32OptionBase struct {
	value float32
}

func (o *Float32OptionBase) Value() interface{} {
	return o.value
}

func (o *Float32OptionBase) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatFloat(float64(o.value), 'f', -1, 32)), nil
}

func (o *Float32OptionBase) UnmarshalText(text []byte) error {
	v, err := strconv.ParseFloat(string(text), 32)
	if err != nil {
		return err
	}
	o.value = float32(v)
	return nil
}

func (o *Float32OptionBase) Default() interface{} {
	return float32(0)
}

var _ Value = (*Float32OptionBase)(nil)
