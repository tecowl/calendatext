package calendatext

import (
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

func (wd Weekday) Match(d *Date) bool {
	if d == nil {
		return false
	}
	return time.Weekday(wd) == d.Time().Weekday()
}

var WeekdayNameMap = map[Weekday]string{
	Sunday:    "日",
	Monday:    "月",
	Tuesday:   "火",
	Wednesday: "水",
	Thursday:  "木",
	Friday:    "金",
	Saturday:  "土",
}
