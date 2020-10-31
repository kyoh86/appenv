package types

import (
	"strconv"
)

type Float64OptionBase struct {
	value float64
}

func (o *Float64OptionBase) Value() interface{} {
	return o.value
}

func (o *Float64OptionBase) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatFloat(float64(o.value), 'f', -1, 64)), nil
}

func (o *Float64OptionBase) UnmarshalText(text []byte) error {
	v, err := strconv.ParseFloat(string(text), 64)
	if err != nil {
		return err
	}
	o.value = float64(v)
	return nil
}

func (o *Float64OptionBase) Default() interface{} {
	return float64(0)
}

var _ Value = (*Float64OptionBase)(nil)
