package calendatext

import (
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// RFC3339 full-time
// See documents about RFC3339
//    https://www.ietf.org/rfc/rfc3339.txt
//    https://medium.com/easyread/understanding-about-rfc-3339-for-datetime-formatting-in-software-engineering-940aa5d5f68a
//    https://wiki.suikawiki.org/n/RFC%203339の日付形式
const DateFormat = "2006-01-02"

type Date struct {
	y int
	m time.Month
	d int
}

func ParseDateWith(str string, delimiter string) (*Date, error) {
	parts := strings.SplitN(str, delimiter, 3)
	if len(parts) < 3 {
		return nil, errors.Errorf("Invalid Date format: %q", str)
	}
	nums := make([]int, 3)
	for idx, s := range parts {
		v, err := strconv.Atoi(s)
		if err != nil {
			return nil, errors.Errorf("Invalid number for date: %q", str)
		}
		nums[idx] = v
	}
	return NewDate(nums[0], time.Month(nums[1]), nums[2]), nil
}

func NewDate(y int, m time.Month, d int) *Date {
	t := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	return &Date{t.Year(), t.Month(), t.Day()}
}

func (d Date) Clone() *Date {
	return &Date{d.y, d.m, d.d}
}

func (d Date) Time() time.Time {
	return time.Date(d.y, d.m, d.d, 0, 0, 0, 0, time.UTC)
}

func (d Date) String() string {
	return d.Time().Format(DateFormat)
}

func (d Date) Equal(other *Date) bool {
	if other == nil {
		return false
	}
	return d.y == other.y && d.m == other.m && d.d == other.d
}

// Implement DateMatcher interface
func (d Date) Match(other *Date) bool {
	return d.Equal(other)
}

func (d Date) beforeAfter(other *Date, compare func(a, b int) bool, resultForSame bool) bool {
	if other == nil {
		return false
	}
	if compare(d.y, other.y) {
		return true
	} else if d.y == other.y {
		if compare(int(d.m), int(other.m)) {
			return true
		} else if d.m == other.m {
			if compare(d.d, other.d) {
				return true
			} else if d.d == other.d {
				return resultForSame
			}
		}
	}
	return false
}

func (d Date) After(other *Date) bool {
	return d.beforeAfter(other, func(a, b int) bool {
		return a > b
	}, false)
}

func (d Date) AfterEqual(other *Date) bool {
	return d.beforeAfter(other, func(a, b int) bool {
		return a > b
	}, true)
}

func (d Date) Before(other *Date) bool {
	return d.beforeAfter(other, func(a, b int) bool {
		return a < b
	}, false)
}

func (d Date) BeforeEqual(other *Date) bool {
	return d.beforeAfter(other, func(a, b int) bool {
		return a < b
	}, true)
}

func (d Date) NextDay() *Date {
	return d.NextDayOf(1)
}

func (d Date) NextWeek() *Date {
	return d.NextWeekOf(1)
}

func (d Date) PrevDay() *Date {
	return d.PrevDayOf(1)
}

func (d Date) PrevWeek() *Date {
	return d.PrevWeekOf(1)
}

func (d Date) NextDayOf(v int) *Date {
	return NewDate(d.y, d.m, d.d+v)
}

func (d Date) NextWeekOf(v int) *Date {
	return d.NextDayOf(v * 7)
}

func (d Date) PrevDayOf(v int) *Date {
	return d.NextDayOf(-1 * v)
}

func (d Date) PrevWeekOf(v int) *Date {
	return d.NextWeekOf(-1 * v)
}

func (d Date) NextMonth() *Date {
	return d.NextMonthOf(1)
}

func (d Date) PrevMonth() *Date {
	return d.PrevMonthOf(1)
}

func (d Date) NextMonthOf(v time.Month) *Date {
	return NewDate(d.y, d.m+v, d.d)
}

func (d Date) PrevMonthOf(v time.Month) *Date {
	return d.NextMonthOf(v * -1)
}

func (d Date) NextYear() *Date {
	return d.NextYearOf(1)
}

func (d Date) PrevYear() *Date {
	return d.PrevYearOf(1)
}

func (d Date) NextYearOf(v int) *Date {
	return NewDate(d.y+v, d.m, d.d)
}

func (d Date) PrevYearOf(v int) *Date {
	return d.NextYearOf(v * -1)
}
