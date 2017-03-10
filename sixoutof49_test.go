package sixoutof49_test

import (
	"github.com/XelaRellum/sixoutof49"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	n, err := sixoutof49.Create()

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
