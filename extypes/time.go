package extypes

import "time"

type TimeOptionBase struct {
	value time.Time
}

func (o *TimeOptionBase) Value() interface{} {
	return o.value
}

func (o *TimeOptionBase) MarshalText() (text []byte, err error) {
	return []byte(o.value.Format(time.RFC3339)), nil
}

func (o *TimeOptionBase) UnmarshalText(text []byte) error {
	t, err := time.Parse(time.RFC3339, string(text))
	if err != nil {
		return err
	}
	o.value = t
	return nil
}

func (o *TimeOptionBase) Default() interface{} {
	return time.Time{}
}
