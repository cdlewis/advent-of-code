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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// https://github.com/golang/go/issues/448
func Mod(d, m int) int {
	res := d % m

	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	
	return res
}
