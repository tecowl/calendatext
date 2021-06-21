package calendatext_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tecowl/calendatext"
)

func TestCalendarWithMonthlyWeekday(t *testing.T) {

	// ----- 2020-12 ------  ----- 2021-01 ------
	//  S  M  T  W  T  F  S    S  M  T  W  T  F  S
	//        1  2  3  4  5                   1  2
	//  6  7  8  9 10 11 12    3  4  5  6  7  8  9
	// 13 14 15 16 17 18 19   10 11 12 13 14 15 16
	// 20 21 22 23 24 25 26   17 18 19 20 21 22 23
	// 27 28 29 30 31         24 25 26 27 28 29 30
	//                       31

	text := `
+ 毎月第1水曜日 : 通常営業日
+ 毎月第3水曜日 : 通常営業日
- 2020/12/28 - 01/06 : 冬休み
-      01/11 : 成人の日
+ 2021/01/07 : 振替営業日
`

	cal := calendatext.NewCalendar(
		*calendatext.NewDate(2020, 12, 1),
		*calendatext.NewDate(2021, 1, 31),
		false,
	)

	cal.ParseText(text)

	assert.Equal(t,
		[]string{
			"2020-12-02",
			"2020-12-16",
			"2021-01-07",
			"2021-01-20",
		},
		cal.Dates().Strings(),
	)

}
