package types

import (
	"strconv"
)

type Float32Value struct {
	value float32
}

func (o *Float32Value) Value() interface{} {
	return o.value
}

func (o *Float32Value) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatFloat(float64(o.value), 'f', -1, 32)), nil
}

func (o *Float32Value) UnmarshalText(text []byte) error {
	v, err := strconv.ParseFloat(string(text), 32)
	if err != nil {
		return err
	}
	o.value = float32(v)
	return nil
}

func (o *Float32Value) Default() interface{} {
	return float32(0)
}

var _ Value = (*Float32Value)(nil)
