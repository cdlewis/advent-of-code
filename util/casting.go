package util

import (
	"fmt"
	"strconv"
)

// Do digusting things to make AoC questions slightly quicker to solve
func ToInt[U any](arg U) int {
	switch any(arg).(type) {
	case byte:
		val, err := strconv.Atoi(string(any(arg).(byte)))
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
		return val
	case rune:
		val, err := strconv.Atoi(string(any(arg).(rune)))
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
		return val
	case []byte:
		val, err := strconv.Atoi(string(any(arg).([]byte)))
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
		return val
	case string:
		val, err := strconv.Atoi(any(arg).(string))
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
		return val
	default:
		panic(fmt.Sprintf("unhandled type for int casting (%T)", arg))
	}
}

func ToString(arg interface{}) string {
	var str string
	switch arg.(type) {
	case int:
		str = strconv.Itoa(arg.(int))
	case byte:
		b := arg.(byte)
		str = string(rune(b))
	case rune:
		str = string(arg.(rune))
	default:
		panic(fmt.Sprintf("unhandled type for string casting %T", arg))
	}
	return str
}
