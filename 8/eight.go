package eight

import (
	"strings"

	"github.com/cdlewis/advent-of-code/util"
)

func Eight(useTest bool, testInput string) int {
	raw := util.Map(strings.Split(util.GetInput(8, useTest, testInput), "\n"), func(line string) []int {
		return util.Map([]byte(line), util.ToInt[byte])
	})

	score := 0
	for i := 0; i < len(raw); i++ {
		for j := 0; j < len(raw[i]); j++ {
			score = util.Max(score, isVisible(i, j, raw))
		}
	}

	return score
}

func isVisible(startI int, startJ int, grid [][]int) int {
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	directionCounts := util.Map(directions, func(d []int) int {
		count := 0
		for i, j := startI+d[0], startJ+d[1]; util.ValidCoordinate(i, j, grid); i, j = i+d[0], j+d[1] {
			count++
			if grid[i][j] >= grid[startI][startJ] {
				break
			}
		}
		return count
	})

	return util.Reduce(directionCounts, util.Multiply, 1)
}
