package calendatext

type Dates []*Date

func (s Dates) Strings() []string {
	r := make([]string, len(s))
	for i, d := range s {
		r[i] = d.String()
	}
	return r
}

func (s Dates) Match(d *Date) bool {
	for _, i := range s {
		if i.Equal(d) {
			return true
		}
	}
	return false
}
