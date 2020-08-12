package calendatext

type Pattern struct {
	Enabled     bool
	Description string
	DateMatcher
}

func NewPattern(enabled bool, description string, matcher DateMatcher) *Pattern {
	return &Pattern{
		Enabled:     enabled,
		Description: description,
		DateMatcher: matcher,
	}
}
