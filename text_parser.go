package calendatext

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

type textParser struct {
	currentDate Date
	Patterns    Patterns
}

func newTextParser(d Date) *textParser {
	return &textParser{currentDate: d, Patterns: Patterns{}}
}

func (tp *textParser) Run(s string) error {
	lines := strings.Split(s, "\n")
	for _, rawLine := range lines {
		line := strings.TrimSpace(rawLine)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		pattern, err := tp.parseLine(line)
		if err != nil {
			return err
		}

		tp.Patterns = append(tp.Patterns, pattern)
	}

	return nil
}

func (tp *textParser) parseLine(line string) (*Pattern, error) {
	var enabled bool
	if strings.HasPrefix(line, "+") {
		enabled = true
	} else if strings.HasPrefix(line, "-") {
		enabled = false
	} else {
		return nil, errors.Errorf("Invalid first charactor. It must be '+' or '-': %q\n", line)
	}

	line = line[1:]
	bodies := strings.SplitN(line, ":", 2)
	description := ""
	if len(bodies) == 2 {
		description = strings.TrimSpace(bodies[1])
	}

	matcher, err := tp.parseMatcher(strings.TrimSpace(bodies[0]))
	if err != nil {
		return nil, errors.WithMessagef(err, "Failed to build matcher for %s", description)
	}
	return &Pattern{
		Enabled:     enabled,
		Description: description,
		DateMatcher: matcher,
	}, nil
}

func (tp *textParser) parseMatcher(body string) (DateMatcher, error) {
	for _, build := range matcherBuilders {
		m, err := build(body)
		if err != nil {
			return nil, err
		}
		if m != nil {
			return m, nil
		}
	}
	return nil, errors.Errorf("No build function found for %q", body)
}

type BuildMatcher func(s string) (DateMatcher, error)

var (
	slashDateRE   = regexp.MustCompile(`\A\d+/\d+/\d+\z`)
	slashPeriodRE = regexp.MustCompile(`\A\d+/\d+/\d+\s*-\s*\d+/\d+/\d+\z`)
)

var matcherBuilders = []BuildMatcher{
	func(s string) (DateMatcher, error) {
		if s != "平日" {
			return nil, nil
		}
		return Weekdays{Monday, Tuesday, Wednesday, Thursday, Friday}, nil
	},

	func(s string) (DateMatcher, error) {
		if !slashDateRE.MatchString(s) {
			return nil, nil
		}
		return ParseDateWith(strings.TrimSpace(s), "/")
	},

	func(s string) (DateMatcher, error) {
		if !slashPeriodRE.MatchString(s) {
			return nil, nil
		}
		parts := strings.SplitN(s, "-", 2)
		if len(parts) < 2 {
			return nil, errors.Errorf("Failed to split string as Period: %q", s)
		}

		st, err := ParseDateWith(strings.TrimSpace(parts[0]), "/")
		if err != nil {
			return nil, err
		}
		ed, err := ParseDateWith(strings.TrimSpace(parts[1]), "/")
		if err != nil {
			return nil, err
		}
		return NewPeriod(*st, *ed), nil
	},
}
