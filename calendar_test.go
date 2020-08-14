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

func TestCalendarParse(t *testing.T) {
	texts := map[string]string{
		"full-date": `
+ 平日 : 通常営業日
- 2020/08/10 : 山の日
- 2020/08/17-2020/08/21 : 夏季休暇
+ 2020/08/22 : 出勤日
`,
		"mmdd": `
+ 平日 : 通常営業日
- 08/10 : 山の日
- 08/17-08/21 : 夏季休暇
+ 08/22 : 出勤日
`,
	}

	for name, text := range texts {
		t.Run(name, func(t *testing.T) {
			c := NewCalendar(*NewDate(2020, 8, 1), *NewDate(2020, 8, 31), false)
			err := c.ParseText(text)
			if !assert.NoError(t, err) {
				return
			}

			if !assert.Equal(t, 4, len(c.Patterns)) {
				return
			}
			if assert.IsType(t, (*Date)(nil), c.Patterns[0].DateMatcher) {
				i := c.Patterns[0]
				m := i.DateMatcher.(*Date)
				assert.True(t, i.Enabled)
				assert.Equal(t, "出勤日", i.Description)
				assert.Equal(t, "2020-08-22", m.String())
			}

			if assert.IsType(t, (*Period)(nil), c.Patterns[1].DateMatcher) {
				i := c.Patterns[1]
				m := i.DateMatcher.(*Period)
				assert.False(t, i.Enabled)
				assert.Equal(t, "夏季休暇", i.Description)
				assert.Equal(t, "2020-08-17", m.Start.String())
				assert.Equal(t, "2020-08-21", m.End.String())
			}

			if assert.IsType(t, (*Date)(nil), c.Patterns[2].DateMatcher) {
				i := c.Patterns[2]
				m := i.DateMatcher.(*Date)
				assert.False(t, i.Enabled)
				assert.Equal(t, "山の日", i.Description)
				assert.Equal(t, "2020-08-10", m.String())
			}

			if assert.IsType(t, (Weekdays)(nil), c.Patterns[3].DateMatcher) {
				i := c.Patterns[3]
				m := i.DateMatcher.(Weekdays)
				assert.True(t, i.Enabled)
				assert.Equal(t, "通常営業日", i.Description)
				assert.Equal(t, Weekdays{Monday, Tuesday, Wednesday, Thursday, Friday}, m)
			}
		})
	}
}
