package calendatext

type Weekdays []Weekday

func (s Weekdays) Match(d *Date) bool {
	dd := Weekday(d.Time().Weekday())
	for _, i := range s {
		if dd == i {
			return true
		}
	}
	return false
}
