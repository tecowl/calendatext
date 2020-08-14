package calendatext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDate(t *testing.T) {
	assert.Equal(t, "2019-11-30", NewDate(2020, 0, 0).String())
	assert.Equal(t, "2020-08-01", NewDate(2020, 8, 1).String())
	assert.Equal(t, "2019-03-01", NewDate(2019, 2, 29).String())
	assert.Equal(t, "2020-03-01", NewDate(2020, 2, 30).String())
	assert.Equal(t, "2028-06-07", NewDate(2020, 99, 99).String())
}

func TestNextDay(t *testing.T) {
	d := NewDate(2020, 8, 12)
	assert.Equal(t, "2020-07-29", d.PrevWeekOf(2).String())
	assert.Equal(t, "2020-08-05", d.PrevWeek().String())
	assert.Equal(t, "2020-08-10", d.PrevDayOf(2).String())
	assert.Equal(t, "2020-08-11", d.PrevDay().String())
	assert.Equal(t, "2020-08-13", d.NextDay().String())
	assert.Equal(t, "2020-08-13", d.NextDayOf(1).String())
	assert.Equal(t, "2020-08-14", d.NextDayOf(2).String())
	assert.Equal(t, "2020-08-19", d.NextWeek().String())
	assert.Equal(t, "2020-09-02", d.NextWeekOf(3).String())
}

func TestNextMonth(t *testing.T) {
	d := NewDate(2020, 8, 12)
	assert.Equal(t, "2019-11-12", d.PrevMonthOf(9).String())
	assert.Equal(t, "2019-12-12", d.PrevMonthOf(8).String())
	assert.Equal(t, "2020-01-12", d.PrevMonthOf(7).String())
	assert.Equal(t, "2020-07-12", d.PrevMonth().String())
	assert.Equal(t, "2020-09-12", d.NextMonth().String())
	assert.Equal(t, "2020-12-12", d.NextMonthOf(4).String())
	assert.Equal(t, "2021-01-12", d.NextMonthOf(5).String())
}

func TestNextYear(t *testing.T) {
	d := NewDate(2020, 8, 12)
	assert.Equal(t, "2018-08-12", d.PrevYearOf(2).String())
	assert.Equal(t, "2019-08-12", d.PrevYear().String())
	assert.Equal(t, "2021-08-12", d.NextYear().String())
	assert.Equal(t, "2023-08-12", d.NextYearOf(3).String())
}

func TestCompare(t *testing.T) {
	t.Run("compare with myself", func(t *testing.T) {
		d := NewDate(2020, 8, 12)
		assert.True(t, d.Equal(d))
		assert.True(t, d.BeforeEqual(d))
		assert.True(t, d.AfterEqual(d))
		assert.False(t, d.Before(d))
		assert.False(t, d.After(d))
	})

	subTestForNext := func(d0, d1 *Date) func(t *testing.T) {
		return func(t *testing.T) {
			assert.False(t, d0.Equal(d1))
			assert.True(t, d0.BeforeEqual(d1))
			assert.True(t, d0.Before(d1))
			assert.False(t, d0.AfterEqual(d1))
			assert.False(t, d0.After(d1))
		}
	}

	subTestForPrev := func(d0, d1 *Date) func(t *testing.T) {
		return func(t *testing.T) {
			assert.False(t, d0.Equal(d1))
			assert.False(t, d0.BeforeEqual(d1))
			assert.False(t, d0.Before(d1))
			assert.True(t, d0.AfterEqual(d1))
			assert.True(t, d0.After(d1))
		}
	}

	d := NewDate(2020, 8, 12)
	t.Run("compare with next day", subTestForNext(d, d.NextDay()))
	t.Run("compare with next month", subTestForNext(d, d.NextMonth()))
	t.Run("compare with next year", subTestForNext(d, d.NextYear()))
	t.Run("compare with prev day", subTestForPrev(d, d.PrevDay()))
	t.Run("compare with prev month", subTestForPrev(d, d.PrevMonth()))
	t.Run("compare with prev year", subTestForPrev(d, d.PrevYear()))
}
