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
	if !iterable.Contains[T](s.items, value) {
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
