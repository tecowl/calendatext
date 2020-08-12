package calendatext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeriodInclude(t *testing.T) {
	t.Run("single date", func(t *testing.T) {
		d := NewDate(2020, 8, 6)
		pd := NewPeriod(*d, *d)
		assert.False(t, pd.Include(d.PrevYear()))
		assert.False(t, pd.Include(d.PrevMonth()))
		assert.False(t, pd.Include(d.PrevWeek()))
		assert.False(t, pd.Include(d.PrevDay()))
		assert.True(t, pd.Include(d))
		assert.False(t, pd.Include(d.NextDay()))
		assert.False(t, pd.Include(d.NextWeek()))
		assert.False(t, pd.Include(d.NextMonth()))
		assert.False(t, pd.Include(d.NextYear()))
	})

	t.Run("normal", func(t *testing.T) {
		d1 := NewDate(2020, 8, 6)
		d2 := NewDate(2020, 8, 20)
		pd := NewPeriod(*d1, *d2)
		assert.False(t, pd.Include(d1.PrevYear()))
		assert.False(t, pd.Include(d1.PrevMonth()))
		assert.False(t, pd.Include(d1.PrevWeek()))
		assert.False(t, pd.Include(d1.PrevDay()))
		assert.True(t, pd.Include(d1))
		assert.True(t, pd.Include(d1.NextDay()))
		assert.True(t, pd.Include(NewDate(2020, 8, 13)))
		assert.True(t, pd.Include(d2.PrevDay()))
		assert.True(t, pd.Include(d2))
		assert.False(t, pd.Include(d2.NextDay()))
		assert.False(t, pd.Include(d2.NextWeek()))
		assert.False(t, pd.Include(d2.NextMonth()))
		assert.False(t, pd.Include(d2.NextYear()))
	})
}
