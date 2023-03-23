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

func TestSet_IsDisjoint(t *testing.T) {
	t.Run("Returns true if set is disjointed", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")
		otherSet := Set[string]{}
		otherSet.Add("world")

		assert.True(t, set.IsDisjoint(otherSet))
	})

	t.Run("Returns false is set isn't disjointed", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")
		otherSet := Set[string]{}
		otherSet.Add("hello")

		assert.False(t, set.IsDisjoint(otherSet))
	})
}

func TestSet_IsSubset(t *testing.T) {
	t.Run("Returns true if set is a subset", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")

		otherSet := Set[string]{}
		otherSet.Add("hello")
		otherSet.Add("world")

		assert.True(t, set.IsSubset(otherSet))
	})

	t.Run("Returns false if set isn't a subset", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")

		otherSet := Set[string]{}
		otherSet.Add("world")

		assert.False(t, set.IsSubset(otherSet))
	})

	t.Run("Returns false if set isn't a true subset", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")

		otherSet := Set[string]{}
		otherSet.Add("hello")

		assert.False(t, set.IsSubset(otherSet))
	})
}

func TestSet_IsSuperset(t *testing.T) {
	t.Run("Returns true if set is a superset", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")
		set.Add("world")

		otherSet := Set[string]{}
		otherSet.Add("world")

		assert.True(t, set.IsSuperset(otherSet))
	})

	t.Run("Returns false if set isn't a superset", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")

		otherSet := Set[string]{}
		otherSet.Add("world")

		assert.False(t, set.IsSuperset(otherSet))
	})

	t.Run("Returns false if set isn't a true superset", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")

		otherSet := Set[string]{}
		otherSet.Add("hello")

		assert.False(t, set.IsSuperset(otherSet))
	})
}

func TestSet_Union(t *testing.T) {
	t.Run("Returns the union of two sets", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")

		otherSet := Set[string]{}
		otherSet.Add("world")

		newSet := set.Union(otherSet)

		assert.Equal(t, []string{"hello", "world"}, newSet.Items())
	})

	t.Run("Returns the union of multiple sets", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")

		otherSet := Set[string]{}
		otherSet.Add("world")

		anotherSet := Set[string]{}
		anotherSet.Add("again")

		newSet := set.Union(otherSet, anotherSet)

		assert.Equal(t, []string{"hello", "world", "again"}, newSet.Items())
	})
}

func TestSet_Intersection(t *testing.T) {
	t.Run("Returns the intersection of two sets", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")
		set.Add("world")

		otherSet := Set[string]{}
		otherSet.Add("hello")
		otherSet.Add("again")

		newSet := set.Intersection(otherSet)

		assert.Equal(t, []string{"hello"}, newSet.Items())
	})

	t.Run("Returns the intersection of multiple sets", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")
		set.Add("world")

		otherSet := Set[string]{}
		otherSet.Add("hello")
		otherSet.Add("world")
		otherSet.Add("again")

		anotherSet := Set[string]{}
		anotherSet.Add("hello")
		anotherSet.Add("world")
		anotherSet.Add("bye")

		newSet := set.Intersection(otherSet, anotherSet)

		assert.Equal(t, []string{"hello", "world"}, newSet.Items())
	})
}

func TestSet_Difference(t *testing.T) {
	t.Run("Returns the difference of two sets", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")
		set.Add("world")

		otherSet := Set[string]{}
		otherSet.Add("bye")
		otherSet.Add("world")

		newSet := set.Difference(otherSet)

		assert.Equal(t, []string{"hello"}, newSet.Items())
	})

	t.Run("Returns the differences of multiple sets", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")
		set.Add("world")

		otherSet := Set[string]{}
		otherSet.Add("bye")
		otherSet.Add("world")

		anotherSet := Set[string]{}
		anotherSet.Add("another")
		anotherSet.Add("world")

		newSet := set.Difference(otherSet, anotherSet)

		assert.Equal(t, []string{"hello"}, newSet.Items())
	})
}

func TestSet_SymmetricDifference(t *testing.T) {
	t.Run("Returns a symmetric difference", func(t *testing.T) {
		set := Set[string]{}
		set.Add("hello")
		set.Add("world")

		otherSet := Set[string]{}
		otherSet.Add("hello")
		otherSet.Add("bye")

		newSet := set.SymmetricDifference(otherSet)

		assert.Equal(t, []string{"world", "bye"}, newSet.Items())
	})
}
