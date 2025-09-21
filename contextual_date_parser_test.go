package calendatext

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContextualDateParser(t *testing.T) {
	t.Parallel()

	cp := NewContextualDateParser("/-", NewDate(2020, 8, 14))

	{
		d, err := cp.Parse("2020/08/15")
		require.NoError(t, err)
		assert.Equal(t, 2020, d.Year())
		assert.Equal(t, time.Month(8), d.Month())
		assert.Equal(t, 15, d.Day())
	}

	{
		d, err := cp.Parse("08/16")
		require.NoError(t, err)
		assert.Equal(t, 2020, d.Year()) // 2020 が補完される
		assert.Equal(t, time.Month(8), d.Month())
		assert.Equal(t, 16, d.Day())
	}

	{
		d, err := cp.Parse("2020/07/16")
		require.NoError(t, err)
		assert.Equal(t, 2020, d.Year()) // 2020 が補完される
		assert.Equal(t, time.Month(7), d.Month())
		assert.Equal(t, 16, d.Day())
	}

	{
		d, err := cp.Parse("17")
		require.NoError(t, err)
		assert.Equal(t, 2020, d.Year())           // 2020 が補完される
		assert.Equal(t, time.Month(7), d.Month()) // 7 が補完される
		assert.Equal(t, 17, d.Day())
	}

	// 同じ日ならば月が進んだりしない
	{
		d, err := cp.Parse("17")
		require.NoError(t, err)
		assert.Equal(t, 2020, d.Year())           // 2020 が補完される
		assert.Equal(t, time.Month(7), d.Month()) // 7 が補完される
		assert.Equal(t, 17, d.Day())
	}

	{
		d, err := cp.Parse("16") // 17よりも前の日なので翌月と判断される
		require.NoError(t, err)
		assert.Equal(t, 2020, d.Year())           // 2020 が補完される
		assert.Equal(t, time.Month(8), d.Month()) // 8 が補完される
		assert.Equal(t, 16, d.Day())
	}

	{
		d, err := cp.Parse("08/15") // 2020/08/15 ではなく 2021/08/15 と解釈される
		require.NoError(t, err)
		assert.Equal(t, 2021, d.Year()) // 2021 が補完される
		assert.Equal(t, time.Month(8), d.Month())
		assert.Equal(t, 15, d.Day())
	}

	{
		d, err := cp.Parse("31")
		require.NoError(t, err)
		assert.Equal(t, 2021, d.Year())           // 2021 が補完される
		assert.Equal(t, time.Month(8), d.Month()) // 8 が補完される
		assert.Equal(t, 31, d.Day())
	}

	{
		d, err := cp.Parse("1")
		require.NoError(t, err)
		assert.Equal(t, 2021, d.Year())           // 2021 が補完される
		assert.Equal(t, time.Month(9), d.Month()) // 9 が補完される
		assert.Equal(t, 1, d.Day())
	}

	{
		d, err := cp.Parse("12/24")
		require.NoError(t, err)
		assert.Equal(t, 2021, d.Year()) // 2021 が補完される
		assert.Equal(t, time.Month(12), d.Month())
		assert.Equal(t, 24, d.Day())
	}

	{
		d, err := cp.Parse("3")
		require.NoError(t, err)
		assert.Equal(t, 2022, d.Year())           // 2022 が補完される
		assert.Equal(t, time.Month(1), d.Month()) // 1 が補完される
		assert.Equal(t, 3, d.Day())
	}
}
