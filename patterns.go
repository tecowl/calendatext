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
