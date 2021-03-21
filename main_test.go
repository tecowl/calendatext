package calendatext_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tecowl/calendatext"
)

// ----- 2020-12 ------  ----- 2021-01 ------
//  S  M  T  W  T  F  S    S  M  T  W  T  F  S
//        1  2  3  4  5                   1  2
//  6  7  8  9 10 11 12    3  4  5  6  7  8  9
// 13 14 15 16 17 18 19   10 11 12 13 14 15 16
// 20 21 22 23 24 25 26   17 18 19 20 21 22 23
// 27 28 29 30 31         24 25 26 27 28 29 30
//                       31

const text = `
+ 平日 : 通常営業日
- 2020/12/28 - 01/06 : 冬休み
-      01/11 : 成人の日
-      01/25-29 : 試験休み
+ 2021/01/04-05 : 出勤日
`

func TestCalendar(t *testing.T) {
	cal := calendatext.NewCalendar(
		*calendatext.NewDate(2020, 12, 1),
		*calendatext.NewDate(2021, 1, 31),
		false,
	)

	cal.ParseText(text)

	assert.Equal(t,
		[]string{
			"2020-12-01",
			"2020-12-02",
			"2020-12-03",
			"2020-12-04",

			"2020-12-07",
			"2020-12-08",
			"2020-12-09",
			"2020-12-10",
			"2020-12-11",

			"2020-12-14",
			"2020-12-15",
			"2020-12-16",
			"2020-12-17",
			"2020-12-18",

			"2020-12-21",
			"2020-12-22",
			"2020-12-23",
			"2020-12-24",
			"2020-12-25",

			"2021-01-04",
			"2021-01-05",
			"2021-01-07",
			"2021-01-08",

			"2021-01-12",
			"2021-01-13",
			"2021-01-14",
			"2021-01-15",

			"2021-01-18",
			"2021-01-19",
			"2021-01-20",
			"2021-01-21",
			"2021-01-22",
		},
		cal.Dates().Strings(),
	)

}

func TestMatch(t *testing.T) {
	parseDate := func(s string) *calendatext.Date {
		r, err := calendatext.ParseDateWith(s, "-")
		assert.NoError(t, err)
		return r
	}
	calendarBase := parseDate("2020-12-01")

	match := func(date string) bool {
		d := parseDate(date)
		r, err := calendatext.Match(text, *calendarBase, *d, false)
		assert.NoError(t, err)
		return r
	}

	assert.False(t, match("2021-01-10"))
	assert.False(t, match("2021-01-11"))
	assert.True(t, match("2021-01-12"))
}
