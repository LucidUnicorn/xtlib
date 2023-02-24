package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet_Add(t *testing.T) {
	t.Run("Adds new element to a set", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")

		assert.Equal(t, []string{"hello"}, set.items)
	})

	t.Run("Doesn't duplicate elements", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")
		set.Add("hello")

		assert.Equal(t, []string{"hello"}, set.items)
	})
}

func TestSet_Remove(t *testing.T) {
	t.Run("Removes an element from the set", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")
		err := set.Remove("hello")

		assert.Nil(t, err)
		assert.Equal(t, []string{}, set.items)
	})

	t.Run("Returns an error if element doesn't exist", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")
		err := set.Remove("invalid")

		assert.Error(t, err)
		assert.Equal(t, []string{"hello"}, set.items)
	})
}

func TestSet_Pop(t *testing.T) {
	t.Run("Removes and returns an element from the set", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")
		item, err := set.Pop("hello")

		assert.Nil(t, err)
		assert.Equal(t, "hello", item)
		assert.Equal(t, []string{}, set.items)
	})

	t.Run("Errors when an item doesn't exist in the set", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")
		item, err := set.Pop("invalid")

		assert.Error(t, err)
		assert.Equal(t, "", item)
		assert.Equal(t, []string{"hello"}, set.items)
	})
}

func TestSet_Discard(t *testing.T) {
	t.Run("Removes an element from the set", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")
		set.Discard("hello")

		assert.Equal(t, []string{}, set.items)
	})

	t.Run("Leaves the set unchanged when element doesn't exist", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")
		set.Discard("invalid")

		assert.Equal(t, []string{"hello"}, set.items)
	})
}

func TestSet_Clear(t *testing.T) {
	t.Run("Empties the set", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")
		set.Add("world")
		set.Clear()

		assert.Empty(t, set.items)
	})
}

func TestSet_Contains(t *testing.T) {
	t.Run("True when set contains element", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")

		assert.True(t, set.Contains("hello"))
	})

	t.Run("False when set doens't contain element", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")

		assert.False(t, set.Contains("invalid"))
	})
}

func TestSet_Items(t *testing.T) {
	t.Run("Returns slice of set elements", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")
		set.Add("world")

		assert.Equal(t, []string{"hello", "world"}, set.Items())
	})
}

func TestSet_Len(t *testing.T) {
	t.Run("Returns current length of set", func(t *testing.T) {
		set := Set[string]{}
		assert.Equal(t, 0, set.Len())

		set.Add("hello")
		set.Add("world")
		assert.Equal(t, 2, set.Len())

		set.Discard("hello")
		assert.Equal(t, 1, set.Len())
	})
}
