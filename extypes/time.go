package extypes

import "time"

type TimePropertyBase struct {
	value time.Time
}

func (o *TimePropertyBase) Value() interface{} {
	return o.value
}

func (o *TimePropertyBase) MarshalText() (text []byte, err error) {
	return []byte(o.value.Format(time.RFC3339)), nil
}

func (o *TimePropertyBase) UnmarshalText(text []byte) error {
	t, err := time.Parse(time.RFC3339, string(text))
	if err != nil {
		return err
	}
	o.value = t
	return nil
}

func (o *TimePropertyBase) Default() interface{} {
	return time.Time{}
}
