package calendatext

type Pattern struct {
	DateMatcher

	Enabled     bool
	Description string
}

func NewPattern(enabled bool, description string, matcher DateMatcher) *Pattern {
	return &Pattern{
		Enabled:     enabled,
		Description: description,
		DateMatcher: matcher,
	}
}
