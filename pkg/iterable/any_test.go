package iterable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAny(t *testing.T) {
	t.Run("Returns true when any value matches target", func(t *testing.T) {
		intVals := []int{0, 0, 1, 0}
		strVals := []string{"a", "a", "b", "a"}
		boolVals := []bool{true, true, false, true}

		assert.True(t, Any[int](intVals, 1))
		assert.True(t, Any[string](strVals, "b"))
		assert.True(t, Any[bool](boolVals, false))
	})
}
