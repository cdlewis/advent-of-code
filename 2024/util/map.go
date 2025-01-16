package util

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
