package calendatext

type DateMatcher interface {
	Match(*Date) bool
}
