package util

import "golang.org/x/exp/constraints"

func Add(x, y int) int { return x + y }

func Multiply(x, y int) int { return x * y }

func Max[U constraints.Ordered](x, y U) U {
	if x > y {
		return x
	}

	return y
}

func Min[U constraints.Ordered](x, y U) U {
	if x < y {
		return x
	}

	return y
}
