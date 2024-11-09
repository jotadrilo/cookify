package slices

func Map[T any, V any](s []T, fn func(T) V) []V {
	var v []V
	for _, item := range s {
		v = append(v, fn(item))
	}
	return v
}

func Select[T any](s []T, fn func(T) bool) []T {
	var v []T
	for _, item := range s {
		if fn(item) {
			v = append(v, item)
		}
	}
	return v
}
