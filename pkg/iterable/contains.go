package iterable

// Contains searches an iterable to determine if it contains a value.
//
//   vals := []int{0, 1, 2, 3, 4}
//
//   if Contains[int](vals, 3) {
//      ...
//   }
func Contains[T comparable](src []T, v T) bool {
	for _, i := range src {
		if i == v {
			return true
		}
	}

	return false
}
