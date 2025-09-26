package calendatext

import (
	"errors"
	"fmt"
	"time"
)

type Weekday time.Weekday

const (
	Sunday    = Weekday(time.Sunday)
	Monday    = Weekday(time.Monday)
	Tuesday   = Weekday(time.Tuesday)
	Wednesday = Weekday(time.Wednesday)
	Thursday  = Weekday(time.Thursday)
	Friday    = Weekday(time.Friday)
	Saturday  = Weekday(time.Saturday)
)

const weekdays = 7

func (wd Weekday) Match(d *Date) bool {
	if d == nil {
		return false
	}
	return time.Weekday(wd) == d.Time().Weekday()
}

// See https://github.com/tecowl/calendatext/issues/5
var WeekdayNameMap = map[Weekday]rune{ // nolint:gochecknoglobals
	Sunday:    '日',
	Monday:    '月',
	Tuesday:   '火',
	Wednesday: '水',
	Thursday:  '木',
	Friday:    '金',
	Saturday:  '土',
}

var ErrUnknownWeekdayName = errors.New("unknown weekday name")

func ParseWeekdayName(s string) (*Weekday, error) {
	c := ([]rune(s))[0]
	for d, name := range WeekdayNameMap {
		if c == name {
			return &d, nil
		}
	}
	return nil, fmt.Errorf("%w %q", ErrUnknownWeekdayName, s)
}
