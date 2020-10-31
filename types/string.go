package types

type StringOptionBase struct {
	value string
}

func (o *StringOptionBase) Value() interface{} {
	return o.value
}

func (o *StringOptionBase) MarshalText() (text []byte, err error) {
	return []byte(o.value), nil
}

func (o *StringOptionBase) UnmarshalText(text []byte) error {
	o.value = string(text)
	return nil
}

func (o *StringOptionBase) Default() interface{} {
	return ""
}

var _ Value = (*StringOptionBase)(nil)
