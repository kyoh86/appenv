package extypes

import "net/url"

type URLOptionBase struct {
	value url.URL
}

func (o *URLOptionBase) Value() interface{} {
	return o.value
}

func (o *URLOptionBase) MarshalText() (text []byte, err error) {
	return []byte(o.value.String()), nil
}

func (o *URLOptionBase) UnmarshalText(text []byte) error {
	u, err := url.Parse(string(text))
	if err != nil {
		return err
	}
	o.value = *u
	return nil
}

func (o *URLOptionBase) Default() interface{} {
	return url.URL{}
}
