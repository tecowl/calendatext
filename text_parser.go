package calendatext

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type BuildMatcher func(s string) (DateMatcher, error)

type textParser struct {
	Patterns        Patterns
	matcherBuilders []BuildMatcher
}

func newTextParser(d Date) *textParser {
	return &textParser{
		Patterns:        Patterns{},
		matcherBuilders: newMatcherBuilders(&d),
	}
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
	switch {
	case strings.HasPrefix(line, "+"):
		enabled = true
	case strings.HasPrefix(line, "-"):
		enabled = false
	default:
		return nil, fmt.Errorf("%w. It must be '+' or '-': %q", ErrInvalidFirstCharacter, line)
	}

	line = line[1:]
	bodies := strings.SplitN(line, ":", 2) // nolint:mnd
	description := ""
	if len(bodies) == 2 { // nolint:mnd
		description = strings.TrimSpace(bodies[1])
	}

	matcher, err := tp.parseMatcher(strings.TrimSpace(bodies[0]))
	if err != nil {
		return nil, fmt.Errorf("%w for %s: %w", ErrMatcherBuild, description, err)
	}
	return &Pattern{
		Enabled:     enabled,
		Description: description,
		DateMatcher: matcher,
	}, nil
}

func (tp *textParser) parseMatcher(body string) (DateMatcher, error) { // nolint:ireturn
	for _, build := range tp.matcherBuilders {
		m, err := build(body)
		if err != nil {
			return nil, err
		}
		if m != nil {
			return m, nil
		}
	}
	return nil, fmt.Errorf("%w for %q", ErrBuildFunctionNotFound, body)
}

var (
	slashDateRE      = regexp.MustCompile(`\A(?:\d+/)?(?:\d+/)?\d+\z`)
	slashPeriodRE    = regexp.MustCompile(`\A(?:\d+/)?(?:\d+/)?\d+\s*-\s*(?:\d+/)?(?:\d+/)?\d+\z`)
	weeklyRE         = regexp.MustCompile(`\A毎週`)
	monthlyDayRE     = regexp.MustCompile(`\A毎月[^\d]*(\d+)日`)
	monthlyWeekdayRE = regexp.MustCompile(`\A毎月.*第(\d)(.+)`)

	ErrInvalidFirstCharacter = errors.New("invalid first character")
	ErrSomethingWrongToParse = errors.New("something wrong to parse")
	ErrPeriodSplit           = errors.New("failed to split string as Period")
	ErrMatcherBuild          = errors.New("failed to build matcher")
	ErrBuildFunctionNotFound = errors.New("no build function found")
)

func newMatcherBuilders(date *Date) []BuildMatcher { // nolint:gocognit,cyclop,funlen
	delimiter := "/"
	contextualParser := NewContextualDateParser(delimiter, date)

	return []BuildMatcher{
		func(s string) (DateMatcher, error) {
			if s != "平日" {
				return nil, nil
			}
			return Weekdays{Monday, Tuesday, Wednesday, Thursday, Friday}, nil
		},

		// 毎週***
		func(s string) (DateMatcher, error) {
			if !weeklyRE.MatchString(s) {
				return nil, nil
			}
			r := Weekdays{}
			for d, c := range WeekdayNameMap {
				if strings.ContainsRune(s, c) {
					r = append(r, d)
				}
			}
			if len(r) == 0 {
				return nil, nil
			}
			return r, nil
		},

		// 毎月***
		func(s string) (DateMatcher, error) {
			m := monthlyDayRE.FindAllStringSubmatch(s, -1)
			if len(m) < 1 {
				return nil, nil
			}
			if len(m[0]) < 2 { // nolint:mnd
				return nil, fmt.Errorf("%w %q", ErrSomethingWrongToParse, s)
			}
			d, err := strconv.ParseInt(m[0][1], 10, 10)
			if err != nil {
				return nil, err // nolint:wrapcheck
			}
			return MonthlyDay(d), nil
		},

		// 毎月第N***
		func(s string) (DateMatcher, error) {
			m := monthlyWeekdayRE.FindAllStringSubmatch(s, -1)
			if len(m) < 1 {
				return nil, nil
			}
			if len(m[0]) < 3 { // nolint:mnd
				return nil, fmt.Errorf("%w %q", ErrSomethingWrongToParse, s)
			}
			n, err := strconv.Atoi(m[0][1])
			if err != nil {
				return nil, err
			}
			wd, err := ParseWeekdayName(m[0][2])
			if err != nil {
				return nil, err
			}
			return &MonthlyWeekday{Num: n, Weekday: *wd}, nil
		},

		func(s string) (DateMatcher, error) {
			if !slashDateRE.MatchString(s) {
				return nil, nil
			}
			return contextualParser.Parse(strings.TrimSpace(s))
		},

		func(s string) (DateMatcher, error) {
			if !slashPeriodRE.MatchString(s) {
				return nil, nil
			}
			parts := strings.SplitN(s, "-", 2) // nolint:mnd
			if len(parts) < 2 {                // nolint:mnd
				return nil, fmt.Errorf("%w %q", ErrPeriodSplit, s)
			}

			st, err := contextualParser.Parse(strings.TrimSpace(parts[0]))
			if err != nil {
				return nil, err
			}
			ed, err := contextualParser.Parse(strings.TrimSpace(parts[1]))
			if err != nil {
				return nil, err
			}
			return NewPeriod(*st, *ed), nil
		},
	}
}
