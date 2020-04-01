package extypes

import "net/url"

type URLPropertyBase struct {
	value url.URL
}

func (o *URLPropertyBase) Value() interface{} {
	return o.value
}

func (o *URLPropertyBase) MarshalText() (text []byte, err error) {
	return []byte(o.value.String()), nil
}

func (o *URLPropertyBase) UnmarshalText(text []byte) error {
	u, err := url.Parse(string(text))
	if err != nil {
		return err
	}
	o.value = *u
	return nil
}

func (o *URLPropertyBase) Default() interface{} {
	return url.URL{}
}
