package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	t.Run("Generates unique string values", func(t *testing.T) {
		r := NewRandomString()
		var strs []string

		for i := 0; i < 1000; i++ {
			strs = append(strs, r.Generate(6))
		}

		var checked []string

		for i := 0; i < len(strs); i++ {
			assert.NotContains(t, checked, strs[i])
			checked = append(checked, strs[i])
		}
	})

	t.Run("Generates with a custom char set", func(t *testing.T) {
		r := String{
			CharSet: []rune("0123456789"),
		}

		for i := 0; i < 100; i++ {
			assert.NotSubset(t, []rune(r.Generate(5)), DefaultStringCharSet)
		}
	})
}
