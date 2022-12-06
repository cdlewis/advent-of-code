package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cdlewis/advent-of-code/util"
)

func parsePairs(input string) [][]int {
	splitInput := strings.FieldsFunc(input, func(r rune) bool { return r == '-' || r == ',' })

	numbers := []int{}
	for _, n := range splitInput {
		result, err := strconv.ParseInt(n, 10, 64)

		if err != nil {
			panic(err)
		}

		numbers = append(numbers, int(result))
	}

	return [][]int{numbers[:2], numbers[2:]}
}

func overlaps(one []int, two []int) bool {
	return (one[0] >= two[0] && one[0] <= two[1]) || (one[1] >= two[0] && one[1] <= two[1])
}

func main() {
	raw := strings.Split(util.GetInput(4, false, ""), "\n")

	result := 0

	for _, data := range raw {
		pairs := parsePairs(data)
		if overlaps(pairs[0], pairs[1]) || overlaps(pairs[1], pairs[0]) {
			result++
		}
	}

	fmt.Println(result)

	if result != 905 {
		panic("unexpected result")
	}
}
