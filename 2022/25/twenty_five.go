package twenty_five

import (
	"strings"

	"github.com/cdlewis/advent-of-code/2022/util"
)

func TwentyFive() string {
	fuelAmounts := util.Map(
		strings.Split(util.GetInput(25, false, ``), "\n"),
		snafuToDecimal,
	)

	totalFuel := util.Reduce(fuelAmounts, util.Add, 0)

	return decimalToSnafu(totalFuel)
}

var decimalDigitToSnafuDigit = map[int]string{
	0: "00",
	1: "01",
	2: "02",
	3: "1=",
	4: "1-",
	5: "10",
	6: "11",
	7: "12",
	8: "2=",
	9: "2-",
}

var addTwoSnafuDigits = map[byte]map[byte][]byte{
	'0': {
		'0': []byte("00"),
		'1': []byte("01"),
		'2': []byte("02"),
		'-': []byte("0-"),
		'=': []byte("0="),
	},
	'1': {
		'0': []byte("00"),
		'1': []byte("02"),
		'2': []byte("1="),
		'-': []byte("00"),
		'=': []byte("0-"),
	},
	'2': {
		'0': []byte("02"),
		'1': []byte("1="),
		'2': []byte("1-"),
		'-': []byte("01"),
		'=': []byte("00"),
	},
	'-': {
		'0': []byte("0-"),
		'1': []byte("00"),
		'2': []byte("01"),
	},
	'=': {
		'0': []byte("0="),
		'1': []byte("0-"),
		'2': []byte("00"),
	},
}

func snafuToDecimal(s string) int {
	power := 0
	result := 0

	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '0', '1', '2':
			result += util.ToInt(s[i]) * util.Pow(5, power)
		case '-':
			result += -1 * util.Pow(5, power)
		case '=':
			result += -2 * util.Pow(5, power)
		}
		power++
	}

	return result
}

func decimalToSnafu(x int) string {
	result := []byte{}

	// Find the required power for the most significant bit

	left, right := 0, 50
	for left < right {
		mid := (left + right) / 2
		powerSize := util.Pow(5, mid)

		if powerSize <= x && util.Pow(5, mid+1) > x {
			break
		}

		if powerSize > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// Progressively subtract powers of five to construct the snafu number

	for pow := left; pow >= 0; pow-- {
		powerSize := util.Pow(5, pow)
		numberOfDivisors := x / powerSize

		// Nothing left to do but add the zeroes

		if x <= 0 || (numberOfDivisors == 0) {
			result = append(result, '0')
			continue
		}

		result = append(result, decimalDigitToSnafuDigit[numberOfDivisors][1])

		carry := decimalDigitToSnafuDigit[numberOfDivisors][0]
		for i := len(result) - 2; i >= 0; i-- {
			if carry == '0' {
				break
			}

			sum := addTwoSnafuDigits[result[i]][byte(carry)]

			if len(sum) > 1 {
				carry = sum[0]
				result[i] = sum[1]
			} else {
				carry = '0'
				result[i] = sum[0]
			}
		}

		if carry != '0' {
			result = append([]byte{carry}, result...)
		}

		x -= (powerSize * numberOfDivisors)
	}

	return string(result)
}
