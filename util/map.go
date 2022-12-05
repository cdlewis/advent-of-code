package util

func MapKeys[T comparable, U any](s map[T]U) []T {
	result := make([]T, len(s))
	
	idx := 0
	for key := range s {
		result[idx] = key
		idx++
	}

	return result
}

func MapValues[T comparable, U any](s map[T]U) []U {
	result := make([]U, len(s))
	
	idx := 0
	for _, value := range s {
		result[idx] = value
		idx++
	}

	return result
}

func MapFromSlice[T comparable](s []T) map[T]bool {
	result := map[T]bool{}
	for _, i := range s {
		result[i] = true
	}
	return result
}

func MapIntersection[T comparable](x, y map[T]bool) map[T]bool {
	result := map[T]bool{}

	for key, present := range x {
		if !present {
			continue
		}

		if _, exists := y[key]; exists {
			result[key] = true
		}
	}

	return result
}
