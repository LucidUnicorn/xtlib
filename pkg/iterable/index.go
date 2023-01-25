package iterable

// Index returns the index of a value within an iterable. If the value doesn't
// exist in the iterable it returns -1.
//
//   vals := []int{0, 1, 2, 3, 4}
//
//   if Index[int](vals, 2) > -1 {
//      ...
//   }
func Index[T comparable](src []T, v T) int {
	for i, x := range src {
		if x == v {
			return i
		}
	}

	return -1
}
