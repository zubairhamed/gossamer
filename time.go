package gossamer

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

var STD_TIME_FORMAT_INSTANT = "2006-01-02T15:04:05-07:00"
var STD_TIME_FORMAT_PERIOD = "2006-01-02T15:04:05.00Z"

func NewDefaultTimePeriod() *TimePeriod {
	return &TimePeriod{}
}

func NewTimePeriod(from time.Time, to time.Time) *TimePeriod {
	return &TimePeriod{
		FromTime: from,
		ToTime:   to,
	}
}

type TimePeriod struct {
	FromTime time.Time
	ToTime   time.Time
}

func (t *TimePeriod) From() time.Time {
	return t.FromTime
}

func (t *TimePeriod) To() time.Time {
	return t.ToTime
}

func (t TimePeriod) GetBSON() (interface{}, error) {
	var out string

	// Time Instant
	if !t.FromTime.IsZero() && !t.ToTime.IsZero() {
		out = fmt.Sprintf("\"%s/%s\"", t.FromTime.Format(STD_TIME_FORMAT_PERIOD), t.ToTime.Format(STD_TIME_FORMAT_PERIOD))
	} else {
		// Time Period
		out = fmt.Sprintf("\"%s\"", t.FromTime.Format(STD_TIME_FORMAT_INSTANT))
	}
	return out, nil
}

func (t *TimePeriod) SetBSON(raw bson.Raw) error {
	var str string
	err := raw.Unmarshal(&str)
	if err == nil {
		str = strings.Replace(str, "\"", "", -1)
		split := strings.Split(str, "/")

		if len(split) == 2 {
			t.FromTime, err = time.Parse(STD_TIME_FORMAT_PERIOD, split[0])
			if err != nil {
				return err
			}
			t.ToTime, err = time.Parse(STD_TIME_FORMAT_PERIOD, split[1])
			if err != nil {
				return err
			}
		} else {
			t.FromTime, err = time.Parse(STD_TIME_FORMAT_INSTANT, split[0])
			if err != nil {
				return err
			}
		}
	}
	return err
}

func (t TimePeriod) MarshalJSON() ([]byte, error) {
	var out string

	// Time Instant
	if !t.FromTime.IsZero() && !t.ToTime.IsZero() {
		out = fmt.Sprintf("\"%s/%s\"", t.From().Format(STD_TIME_FORMAT_PERIOD), t.To().Format(STD_TIME_FORMAT_PERIOD))
	} else {
		// Time Period
		out = fmt.Sprintf("\"%s\"", t.FromTime.Format(STD_TIME_FORMAT_INSTANT))
	}
	return []byte(out), nil
}

func (t *TimePeriod) UnmarshalJSON(data []byte) error {
	var err error
	str := strings.Replace(string(data), "\"", "", -1)
	split := strings.Split(str, "/")

	if len(split) == 2 {
		t.FromTime, err = time.Parse(STD_TIME_FORMAT_PERIOD, split[0])
		if err != nil {
			return err
		}
		t.ToTime, err = time.Parse(STD_TIME_FORMAT_PERIOD, split[1])
		if err != nil {
			return err
		}
	} else {
		t.FromTime, err = time.Parse(STD_TIME_FORMAT_INSTANT, split[0])
		if err != nil {
			return err
		}
	}
	return nil
}

func (t TimePeriod) IsZero() bool {
	return t.FromTime.IsZero() && t.ToTime.IsZero()
}

type TimeInstant time.Time

func NewTimeInstant(t time.Time) *TimeInstant {
	ti := TimeInstant(t)

	return &ti
}

func (t TimeInstant) GetBSON() (interface{}, error) {
	out := fmt.Sprintf("\"%s\"", time.Time(t).Format(STD_TIME_FORMAT_INSTANT))

	return out, nil
}

func (t *TimeInstant) SetBSON(raw bson.Raw) error {
	var str string
	err := raw.Unmarshal(&str)
	if err == nil {
		tv, err := time.Parse(STD_TIME_FORMAT_INSTANT, strings.Replace(str, "\"", "", -1))
		if err != nil {
			return err
		}
		*t = TimeInstant(tv)
	}
	return err
}

func (t TimeInstant) MarshalJSON() ([]byte, error) {
	out := fmt.Sprintf("\"%s\"", time.Time(t).Format(STD_TIME_FORMAT_INSTANT))

	return []byte(out), nil
}

func (t *TimeInstant) UnmarshalJSON(data []byte) error {
	tv, err := time.Parse(STD_TIME_FORMAT_INSTANT, strings.Replace(string(data), "\"", "", -1))
	if err != nil {
		return err
	}

	*t = TimeInstant(tv)

	return nil
}

func (t TimeInstant) IsZero() bool {
	return time.Time(t).IsZero()
}
