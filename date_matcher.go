package calendatext

type DateMatcher interface {
	Match(d *Date) bool
}
