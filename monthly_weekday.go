package calendatext

type MonthlyWeekday struct {
	Num     int
	Weekday Weekday
}

func (mw MonthlyWeekday) Match(d *Date) bool {
	if d == nil {
		return false
	}
	return d.Weekday() == mw.Weekday && d.MonthlyWeekNum() == mw.Num
}
