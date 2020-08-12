package calendatext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalendarDays(t *testing.T) {

	// ----- 2020-08 ------
	// S  M  T  W  T  F  S
	//                    1
	//  2  3  4  5  6  7  8
	//  9 10 11 12 13 14 15
	// 16 17 18 19 20 21 22
	// 23 24 25 26 27 28 29
	// 30 31

	t.Run("single date", func(t *testing.T) {
		c := &Calendar{
			Period:      NewPeriod(*NewDate(2020, 8, 1), *NewDate(2020, 8, 31)),
			BaseEnabled: false,
			Patterns: Patterns{
				NewPattern(true, "出勤日", NewDate(2020, 8, 22)),
				NewPattern(false, "夏季休暇", NewPeriod(*NewDate(2020, 8, 17), *NewDate(2020, 8, 21))),
				NewPattern(false, "山の日", NewDate(2020, 8, 10)),
				NewPattern(true, "平日", Weekdays{
					Monday, Tuesday, Wednesday, Thursday, Friday,
				}),
			},
		}
		assert.Equal(t,
			[]string{
				"2020-08-03",
				"2020-08-04",
				"2020-08-05",
				"2020-08-06",
				"2020-08-07",
				"2020-08-11",
				"2020-08-12",
				"2020-08-13",
				"2020-08-14",
				"2020-08-22",
				"2020-08-24",
				"2020-08-25",
				"2020-08-26",
				"2020-08-27",
				"2020-08-28",
				"2020-08-31",
			},
			c.Dates().Strings(),
		)
	})
}
