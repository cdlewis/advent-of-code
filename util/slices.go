package util

func Map[T, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Reverse[T any](s []T) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

func Flatten[T any](s [][]T) []T {
	result := []T{}

	for _, i := range s {
		result = append(result, i...)
	}

	return result
}

func Intersection[T comparable](slices ...[]T) []T {
	seen := map[T][]bool{}

	for idx, slice := range slices {
		for _, item := range slice {
			if len(seen[item]) == 0 {
				seen[item] = make([]bool, len(slices))
			}

			seen[item][idx] = true
		}
	}

	results := []T{}
Outer:
	for key, timesSeen := range seen {
		for _, wasSeen := range timesSeen {
			if !wasSeen {
				continue Outer
			}
		}

		results = append(results, key)
	}

	return results
}

func IntersectionString(strings ...string) string {
	byteArrays := Map(strings, func(s string) []byte { return []byte(s) })
	result := Intersection(byteArrays...)
	return string(result)
}
