package one

import (
	"strings"

	"github.com/cdlewis/advent-of-code/2024/util"
	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/cast"
)

func One() int {
	pairs := strings.Split(aoc.GetInput(1, false, ""), "\n")

	var leftList []int
	rightListFreq := map[int]int{}
	for _, i := range pairs {
		tokens := strings.Split(i, "   ")
		leftList = append(leftList, cast.ToInt(tokens[0]))
		rightListFreq[cast.ToInt(tokens[1])]++
	}

	totalDist := 0
	for _, i := range leftList {
		totalDist += util.Abs(i * rightListFreq[i])
	}

	return totalDist
}
