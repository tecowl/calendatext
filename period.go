package calendatext

type Period struct {
	Start Date
	End   Date
}

func NewPeriod(start, end Date) *Period {
	return &Period{start, end}
}

func (pd *Period) Include(d *Date) bool {
	if d == nil {
		return false
	}
	return pd.Start.BeforeEqual(d) && pd.End.AfterEqual(d)
}

// Implement DateMatcher interface
func (pd Period) Match(other *Date) bool {
	return pd.Include(other)
}

func (pd *Period) Each(f func(*Date)) {
	curr := pd.Start
	for {
		f(&curr)
		curr = *curr.NextDay()
		if curr.After(&pd.End) {
			break
		}
	}
}
