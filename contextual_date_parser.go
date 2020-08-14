package calendatext

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type ContextualDateParser struct {
	delimiter *regexp.Regexp
	current   *Date
}

func NewContextualDateParser(delimiterPattern string, d *Date) *ContextualDateParser {
	if d == nil {
		d = Today()
	}
	return &ContextualDateParser{
		delimiter: regexp.MustCompile("[" + delimiterPattern + "]"),
		current:   d,
	}
}

func (cp *ContextualDateParser) Parse(s string) (*Date, error) {
	parts := cp.delimiter.Split(strings.TrimSpace(s), 3)
	var y, d int
	var m time.Month
	var err error
	switch len(parts) {
	case 1:
		curr := cp.current
		d, err = strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		if d < curr.Day() {
			m = curr.Month() + 1
		} else {
			m = curr.Month()
		}
		y = curr.Year()
	case 2:
		curr := cp.current
		v, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		m = time.Month(v)
		d, err = strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		tmpD := NewDate(curr.Year(), m, d)
		if curr.After(tmpD) {
			y = curr.Year() + 1
		} else {
			y = curr.Year()
		}
	case 3:
		y, err = strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		v, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		m = time.Month(v)
		d, err = strconv.Atoi(parts[2])
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.Errorf("Something wrong to parse %v (len: %d) as Date", parts, len(parts))
	}

	r := NewDate(y, m, d)
	cp.current = r
	return r, nil
}
