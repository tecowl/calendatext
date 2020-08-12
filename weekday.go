package calendatext

import (
	"time"
)

type Weekday time.Weekday

const (
	Sumday    = Weekday(time.Sunday)
	Monday    = Weekday(time.Monday)
	Tuesday   = Weekday(time.Tuesday)
	Wednesday = Weekday(time.Wednesday)
	Thursday  = Weekday(time.Thursday)
	Friday    = Weekday(time.Friday)
	Saturday  = Weekday(time.Saturday)
)

func (wd Weekday) Match(d *Date) bool {
	if d == nil {
		return false
	}
	return time.Weekday(wd) == d.Time().Weekday()
}
