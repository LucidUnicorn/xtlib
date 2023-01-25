package iterable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAll(t *testing.T) {
	t.Run("Returns true when all values match target", func(t *testing.T) {
		intVals := []int{0, 0, 0, 0, 0}
		strVals := []string{"a", "a", "a", "a"}
		boolVals := []bool{true, true, true}

		assert.True(t, All[int](intVals, 0))
		assert.True(t, All[string](strVals, "a"))
		assert.True(t, All[bool](boolVals, true))
	})

	t.Run("Returns false when any value doesn't match target", func(t *testing.T) {
		intVals := []int{0, 0, 1, 0, 0}
		strVals := []string{"a", "a", "b", "a"}
		boolVals := []bool{true, true, false, true}

		assert.False(t, All[int](intVals, 0))
		assert.False(t, All[string](strVals, "a"))
		assert.False(t, All[bool](boolVals, true))
	})
}
