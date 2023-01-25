package iterable

func All[T comparable](src []T, v T) bool {
	for _, s := range src {
		if s != v {
			return false
		}
	}

	return true
}
