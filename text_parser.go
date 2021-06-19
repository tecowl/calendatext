package calendatext

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
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
	for _, build := range tp.matcherBuilders {
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

var (
	slashDateRE   = regexp.MustCompile(`\A(?:\d+/)?(?:\d+/)?\d+\z`)
	slashPeriodRE = regexp.MustCompile(`\A(?:\d+/)?(?:\d+/)?\d+\s*-\s*(?:\d+/)?(?:\d+/)?\d+\z`)
	weeklyRE      = regexp.MustCompile(`\A毎週`)
	monthlyDayRE  = regexp.MustCompile(`\A毎月(\d+)日`)
	monthlyWeekdayRE = regexp.MustCompile(`\A毎月.*第(\d)(.+)`)
)

func newMatcherBuilders(date *Date) []BuildMatcher {
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
			if len(m[0]) < 2 {
				return nil, errors.Errorf("something wrong to parse %q", s)
			}
			d, err := strconv.ParseInt(m[0][1], 10, 10)
			if err != nil {
				return nil, err
			}
			return MonthlyDay(d), nil
		},

		// 毎月第N***
		func(s string) (DateMatcher, error) {
			m := monthlyWeekdayRE.FindAllStringSubmatch(s, -1)
			if len(m) < 1 {
				return nil, nil
			}
			if len(m[0]) < 3 {
				return nil, errors.Errorf("something wrong to parse %q", s)
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
			parts := strings.SplitN(s, "-", 2)
			if len(parts) < 2 {
				return nil, errors.Errorf("Failed to split string as Period: %q", s)
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
