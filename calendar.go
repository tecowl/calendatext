package calendatext

type Calendar struct {
	Period      *Period
	BaseEnabled bool
	Patterns    Patterns
}

func NewCalendar(start, end Date, baseEnabled bool) *Calendar {
	return &Calendar{Period: NewPeriod(start, end), BaseEnabled: baseEnabled}
}

func (c *Calendar) Dates() Dates {
	r := Dates{}
	c.Period.Each(func(d *Date) {
		x := c.Patterns.FirstMatchAt(d)
		if x == nil {
			if c.BaseEnabled {
				r = append(r, d.Clone())
			}
		} else {
			if x.Enabled {
				r = append(r, d.Clone())
			}
		}
	})
	return r
}
