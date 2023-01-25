package iterable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndex(t *testing.T) {
	t.Run("Returns >-1 when value found", func(t *testing.T) {
		intVals := []int{0, 1, 2, 3, 4}
		strVals := []string{"a", "b", "c", "d"}

		assert.Equal(t, Index[int](intVals, 2), 2)
		assert.Equal(t, Index[string](strVals, "d"), 3)
	})

	t.Run("Returns -1 when value not found", func(t *testing.T) {
		intVals := []int{0, 1, 2, 3, 4}
		strVals := []string{"a", "b", "c", "d"}

		assert.Equal(t, Index[int](intVals, 5), -1)
		assert.Equal(t, Index[string](strVals, "f"), -1)
	})
}
