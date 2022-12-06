package util

import "fmt"

func Map[T, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Pop[T any](s *[]T) T {
	if len(*s) == 0 {
		panic("tried to pop empty array")
	}

	current := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return current
}

func Pops[T any](s *[]T, count int) []T {
	if len(*s) < count {
		panic(fmt.Sprintf("Tried to pop %v items when only %v were present on %v", count, len(*s), *s))
	}

	if count < 0 {
		panic(fmt.Sprintf("Count (%v) is < 0. Cannot pop", count))
	}

	currents := (*s)[len(*s)-count:]
	*s = (*s)[:len(*s)-count]

	return currents
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
