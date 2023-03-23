package set

import (
	"errors"
	"github.com/lucidunicorn/xtlib/pkg/iterable"
)

// Set is an array-like type that can't contain duplicate values.
type Set[T comparable] struct {
	items []T
}

// Add inserts a new item into the set if it doesn't already exist.
//
//    set := Set[string]{}
//    set.Add("hello")
func (s *Set[T]) Add(value T) {
	if !s.Contains(value) {
		s.items = append(s.items, value)
	}
}

// Remove deletes an item from the set. An error is returned if the item
// doesn't exist within the set.
//
//    set := Set[string]{}
//    set.Add("hello")
//    set.Add("world")
//
//    _ := set.Remove("hello")
//
//    if err := set.Remove("invalid"); err != nil {
//        log.Fatal(err)
//    }
func (s *Set[T]) Remove(value T) error {
	for i, item := range s.items {
		if item == value {
			s.items = append(s.items[:i], s.items[i+1:]...)
			return nil
		}
	}

	return errors.New("value is not present in the set")
}

// Pop deletes an item from the set and returns it.
//
//    set := Set[string]{}
//    set.Add("hello")
//    set.Add("world")
//
//    item, err := set.Pop("hello")
//    if err != nil {
//        fmt.Println(item)
//    }
func (s *Set[T]) Pop(value T) (T, error) {
	for i, item := range s.items {
		if item == value {
			s.items = append(s.items[:i], s.items[i+1:]...)
			return item, nil
		}
	}

	return *new(T), errors.New("value is not present in the set")
}

// Discard deletes an item from the set is it exists.
//
//    set := Set[string]{}
//    set.Add("hello")
//    set.Add("world")
//
//    set.Discard("hello")
//    set.Discard("invalid")
func (s *Set[T]) Discard(value T) {
	for i, item := range s.items {
		if item == value {
			s.items = append(s.items[:i], s.items[i+1:]...)
			return
		}
	}
}

// Clear empties the set.
//
//    set := Set[string]{}
//    set.Add("hello")
//    fmt.Println(set.Len()) // output: 1
//
//    set.Clear()
//    fmt.Println(set.Len()) // output: 0
func (s *Set[T]) Clear() {
	s.items = *new([]T)
}

// Contains determines if a given value exists within the set.
//
//    set := Set[string]{}
//    set.Add("hello")
//
//    fmt.Println(set.Contains("hello")) // output: true
//    fmt.Println(set.Contains("invalid")) // output: false
func (s *Set[T]) Contains(value T) bool {
	return iterable.Contains[T](s.items, value)
}

// Items returns the contains of the set as a slice.
//
//    set := Set[string]{}
//    set.Add("hello")
//    set.Add("world")
//
//    for _, item := range set.Items() {
//        fmt.Println(item)
//    }
func (s *Set[T]) Items() []T {
	return s.items
}

// Len returns the number of items in the slice.
//
//    set := Set[string]{}
//
//    fmt.Println(set.Len()) // output: 0
//
//    set.Add("hello")
//    fmt.Println(set.Len()) // output: 1
func (s *Set[T]) Len() int {
	return len(s.items)
}

// IsDisjoint returns true if the set has no elements in common with other.
//
//    set := Set[string]{}
//    set.Add("item1")
//    set.Add("item2")
//
//    otherSet := Set[string]{}
//    otherSet.Add("item3")
//    otherSet.Add("item4")
//    fmt.Println(set.IsDisjoint(otherSet)) // output: true
func (s *Set[T]) IsDisjoint(other Set[T]) bool {
	for _, item := range s.Items() {
		if other.Contains(item) {
			return false
		}
	}

	return true
}

// IsSubset returns true if every item in the set is also in other and other is
// not equal to the set.
//
//    set := Set[string]{}
//    set.Add("item1")
//    set.Add("item2")
//
//    otherSet := Set[string]{}
//    otherSet.Add("item0")
//    fmt.Println(set.IsSubset(otherSet)) // output: false
//
//    otherSet.Add("item1")
//    otherSet.Add("item2")
//    fmt.Println(set.IsSubset(otherSet)) // output: true
func (s *Set[T]) IsSubset(other Set[T]) bool {
	for _, item := range s.Items() {
		if !other.Contains(item) {
			return false
		}
	}

	return s.Len() != other.Len()
}

// IsSuperset returns true if every item in other is also in the set and other
// is not equal to the set.
//
//    set := Set[string]{}
//    set.Add("item0")
//
//    otherSet := Set[string]{}
//    otherSet.Add("item1")
//    otherSet.Add("item2")
//    fmt.Println(set.IsSuperset(otherSet)) // output: false
//
//    set.Add("item1")
//    set.Add("item2")
//    fmt.Println(set.IsSuperset(otherSet)) // output: true
func (s *Set[T]) IsSuperset(other Set[T]) bool {
	for _, item := range other.Items() {
		if !s.Contains(item) {
			return false
		}
	}

	return s.Len() != other.Len()
}

// Union returns a new set containing the elements of the original set combined
// with the elements of all additional sets in others.
//
//    set := Set[string]{}
//    set.Add("item1")
//    otherSet := Set[string]{}
//    otherSet.Add("item2")
//    anotherSet := Set[string]{}
//    anotherSet.Add("item3")
//
//    newSet := set.Union(otherSet, anotherSet)
//    fmt.Println(newSet.Items()) // output: item1, item2, item3
func (s *Set[T]) Union(others ...Set[T]) Set[T] {
	newSet := Set[T]{}

	for _, item := range s.items {
		newSet.Add(item)
	}

	for _, set := range others {
		for _, item := range set.Items() {
			newSet.Add(item)
		}
	}

	return newSet
}

// Intersection returns a new set containing elements that are common to
// the set and all others.
//
//    set := Set[string]{}
//    set.Add("item1")
//    set.Add("item2")
//    otherSet := Set[string]{}
//    otherSet.Add("item1")
//    anotherSet := Set[string]{}
//    anotherSet.Add("item2")
//
//    newSet := set.Intersection(otherSet, anotherSet)
//    fmt.Println(newSet.Items()) // output: item1
func (s *Set[T]) Intersection(others ...Set[T]) Set[T] {
	newSet := Set[T]{}

	for _, item := range s.items {
		foundCount := 0

		for _, set := range others {
			if set.Contains(item) {
				foundCount++
			} else {
				break
			}
		}

		if foundCount == len(others) {
			newSet.Add(item)
		}
	}

	return newSet
}

// Difference returns a new set containing items from the set that don't appear
// in others.
//
//    set := Set[string]{}
//    set.Add("item1")
//    set.Add("item2")
//    set.Add("item3")
//    otherSet := Set[string]{}
//    otherSet.Add("item3")
//    anotherSet := Set[string]{}
//    anotherSet.Add("item3")
//
//    newSet := set.Difference(otherSet, anotherSet)
//    fmt.Println(newSet.Items()) // output: item1, item2
func (s *Set[T]) Difference(others ...Set[T]) Set[T] {
	newSet := Set[T]{}

	for _, item := range s.items {
		found := false

		for _, set := range others {
			if set.Contains(item) {
				found = true
				break
			}
		}

		if !found {
			newSet.Add(item)
		}
	}

	return newSet
}

// SymmetricDifference returns a new set containing items that appear in the
// set or other, but not both.
//
//    set := Set[string]{}
//    set.Add("item1")
//    set.Add("item2")
//    otherSet := Set[string]{}
//    otherSet.Add("item1")
//    otherSet.Add("item3")
//
//    newSet := set.SymmetricDifference(otherSet)
//    fmt.Println(newSet.Items()) // output: item2, item3
func (s *Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	newSet := Set[T]{}
	setDiff := s.Difference(other)
	otherDiff := other.Difference(*s)

	for _, item := range setDiff.Items() {
		newSet.Add(item)
	}

	for _, item := range otherDiff.Items() {
		newSet.Add(item)
	}

	return newSet
}
