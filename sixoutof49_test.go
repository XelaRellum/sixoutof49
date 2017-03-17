package sixoutof49_test

import (
	"github.com/XelaRellum/sixoutof49"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	n, err := sixoutof49.Create(1, 49, 6)

	assert.Nil(t, err, "no error expected")

	t.Run("length", func(t *testing.T) {
		assert.Equal(t, 6, len(n), "should contain exactly 6 numbers")
	})

	t.Run("ascending and unique numbers", func(t *testing.T) {
		last := 0
		for _, value := range n {
			assert.True(t, value > last, "expected greated value")
			last = value
		}
	})

	t.Run("numbers in range 1 to 49", func(t *testing.T) {
		for _, value := range n {
			assert.True(t, value >= 1, "expected >= 1")
			assert.True(t, value <= 49, "expected <= 49")
		}
	})
}

func TestCreateWithIllegalCount(t *testing.T) {
	n, err := sixoutof49.Create(1, 49, 0)

	assert.Nil(t, n, "no numbers expected")
	assert.Equal(t, sixoutof49.ErrCountTooLow, err)
}

func TestCreateWithIllegalMinMax(t *testing.T) {
	n, err := sixoutof49.Create(49, 1, 6)

	assert.Nil(t, n, "no numbers expected")
	assert.Equal(t, sixoutof49.ErrMinGreaterMax, err)
}

func TestCreateWithIllegalMin(t *testing.T) {
	n, err := sixoutof49.Create(0, 49, 6)

	assert.Nil(t, n, "no numbers expected")
	assert.Equal(t, sixoutof49.ErrMinTooLow, err)
}

func TestCreateWithIllegalMax(t *testing.T) {
	n, err := sixoutof49.Create(1, 0, 6)

	assert.Nil(t, n, "no numbers expected")
	assert.Equal(t, sixoutof49.ErrMaxTooLow, err)
}

func TestCreateWithDifferentCount(t *testing.T) {
	n, err := sixoutof49.Create(1, 49, 7)

	assert.Nil(t, err, "no error expected")
	assert.Equal(t, 7, len(n))
}

const (
	MAX_NUMBER_RUNS_TO_FIND = 1000
)

func findNumber(t *testing.T, min int, max int, find int) {
	for i := 0; i < MAX_NUMBER_RUNS_TO_FIND; i++ {
		numbers, err := sixoutof49.Create(min, max, 6)
		assert.Nil(t, err, "no error expected")

		for _, number := range numbers {
			assert.True(t, number >= min, "number greater min")
			assert.True(t, number <= max, "number less max")
			if number == find {
				return // found min
			}
		}
	}

	assert.Fail(t, "value %v not found within %v runs", find, MAX_NUMBER_RUNS_TO_FIND)
}

func TestMinIsFound(t *testing.T) {
	findNumber(t, 1, 49, 1)
}

func TestMaxIsFound(t *testing.T) {
	findNumber(t, 1, 49, 49)
}
