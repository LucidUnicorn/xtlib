package iterable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	t.Run("Returns true when value found", func(t *testing.T) {
		intVals := []int{0, 1, 2, 3, 4}
		strVals := []string{"a", "b", "c", "d"}

		assert.True(t, Contains[int](intVals, 1))
		assert.True(t, Contains[string](strVals, "c"))
	})

	t.Run("Returns false when value not found", func(t *testing.T) {
		intVals := []int{0, 1, 2, 3, 4}
		strVals := []string{"a", "b", "c", "d"}

		assert.False(t, Contains[int](intVals, 5))
		assert.False(t, Contains[string](strVals, "e"))
	})
}
