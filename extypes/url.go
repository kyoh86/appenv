package extypes

import "net/url"

type URLValue struct {
	value url.URL
}

func (o *URLValue) Value() interface{} {
	return o.value
}

func (o *URLValue) MarshalText() (text []byte, err error) {
	return []byte(o.value.String()), nil
}

func (o *URLValue) UnmarshalText(text []byte) error {
	u, err := url.Parse(string(text))
	if err != nil {
		return err
	}
	o.value = *u
	return nil
}

func (o *URLValue) Default() interface{} {
	return url.URL{}
}
