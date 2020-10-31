package types

import (
	"strconv"
)

type Float64Value struct {
	value float64
}

func (o *Float64Value) Value() interface{} {
	return o.value
}

func (o *Float64Value) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatFloat(float64(o.value), 'f', -1, 64)), nil
}

func (o *Float64Value) UnmarshalText(text []byte) error {
	v, err := strconv.ParseFloat(string(text), 64)
	if err != nil {
		return err
	}
	o.value = float64(v)
	return nil
}

func (o *Float64Value) Default() interface{} {
	return float64(0)
}

var _ Value = (*Float64Value)(nil)
