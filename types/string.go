package types

type StringValue struct {
	value string
}

func (o *StringValue) Value() interface{} {
	return o.value
}

func (o *StringValue) MarshalText() (text []byte, err error) {
	return []byte(o.value), nil
}

func (o *StringValue) UnmarshalText(text []byte) error {
	o.value = string(text)
	return nil
}

func (o *StringValue) Default() interface{} {
	return ""
}

var _ Value = (*StringValue)(nil)
