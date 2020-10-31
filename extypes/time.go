package extypes

import "time"

type TimeValue struct {
	value time.Time
}

func (o *TimeValue) Value() interface{} {
	return o.value
}

func (o *TimeValue) MarshalText() (text []byte, err error) {
	return []byte(o.value.Format(time.RFC3339)), nil
}

func (o *TimeValue) UnmarshalText(text []byte) error {
	t, err := time.Parse(time.RFC3339, string(text))
	if err != nil {
		return err
	}
	o.value = t
	return nil
}

func (o *TimeValue) Default() interface{} {
	return time.Time{}
}
