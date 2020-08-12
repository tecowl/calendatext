package calendatext

type Patterns []*Pattern

func (s Patterns) FirstMatchAt(d *Date) *Pattern {
	for _, i := range s {
		if i.Match(d) {
			return i
		}
	}
	return nil
}

func (s Patterns) Reverse() Patterns {
	length := len(s)
	r := make(Patterns, length)
	for idx, i := range s {
		r[length-idx-1] = i
	}
	return r
}
