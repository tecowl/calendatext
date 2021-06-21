package calendatext

type MonthlyDay int

func (md MonthlyDay) Match(d *Date) bool {
	if d == nil {
		return false
	}
	return int(md) == d.Day()
}
