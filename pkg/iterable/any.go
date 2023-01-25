package iterable

func Any[T comparable](src []T, v T) bool {
	for _, s := range src {
		if s == v {
			return true
		}
	}

	return false
}
