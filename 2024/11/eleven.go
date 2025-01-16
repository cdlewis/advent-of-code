package eleven

import (
	"strconv"
	"strings"

	"github.com/cdlewis/advent-of-code/2024/util"
	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/cast"
)

var cache = map[[3]int]int{}

func Eleven() int {
	stones := util.Map(strings.Split(aoc.GetInput(11, false, "125 17"), " "), cast.ToInt)
	return simulate(stones, 75)
}

func simulate(stones []int, steps int) int {
	result := 0

	if len(stones) == 2 {
		key := [3]int{stones[0], stones[1], steps}
		if val, ok := cache[key]; ok {
			return val
		}
		defer func() {
			cache[key] = result
		}()
	}

	for currentStep := range steps {
		for idx, s := range stones {
			if s == -1 {
				continue
			}

			if s == 0 {
				stones[idx] = 1
				continue
			}

			digits := strconv.Itoa(s)
			if len(digits)%2 == 0 {
				stones[idx] = -1

				simResult := simulate(
					[]int{
						cast.ToInt(digits[:len(digits)/2]),
						cast.ToInt(digits[len(digits)/2:]),
					},
					steps-currentStep-1,
				)

				result += simResult
				continue
			}

			stones[idx] *= 2024
		}
	}

	for _, s := range stones {
		if s != -1 {
			result++
		}
	}
	return result
}
