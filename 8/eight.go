package eight

import (
	"strings"

	"github.com/cdlewis/advent-of-code/util"
)

func Eight(useTest bool, testInput string) int {
	raw := strings.Split(util.GetInput(8, useTest, testInput), "\n")

	score := 0
	for i := 0; i < len(raw); i++ {
		for j := 0; j < len(raw[i]); j++ {
			newScore := isVisible(i, j, raw)
			if newScore > score {
				score = newScore
			}
		}
	}

	return score
}

func isVisible(startI int, startJ int, grid []string) int {
	visibleHeight := util.ToInt(grid[startI][startJ])
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	directionCounts := util.Map(directions, func(d []int) int {
		count := 0
		for i, j := startI+d[0], startJ+d[1]; i >= 0 && j >= 0 && i < len(grid) && j < len(grid[i]); i, j = i+d[0], j+d[1] {
			count++
			if util.ToInt(grid[i][j]) >= visibleHeight {
				break
			}
		}
		return count
	})

	return util.Reduce(directionCounts, func(current int, next int) int { return current * next }, 1)
}
