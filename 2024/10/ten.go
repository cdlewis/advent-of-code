package ten

import (
	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/grid"
	"github.com/cdlewis/advent-of-code/2024/util/set"
)

func Ten() int {
	grid := grid.ToGrid(aoc.GetInput(10, false, ""))

	score := 0
	for idx, i := range grid {
		for jdx, j := range i {
			if j == 0 {
				score += countTrails([2]int{idx, jdx}, grid, set.New[[2]int]())
			}
		}
	}

	return score
}

func countTrails(start grid.Point, trailMap grid.Grid[int], seen set.Set[[2]int]) int {
	if seen.Exists(start) {
		return 0
	}

	if trailMap.Get(start) == 9 {
		return 1
	}

	seen.Add(start)

	score := 0
	for _, newPosition := range trailMap.GetAdjacent(start) {
		if trailMap.Get(newPosition) != trailMap.Get(start)+1 {
			continue
		}

		score += countTrails(newPosition, trailMap, seen)
	}

	seen.Remove(start)

	return score
}
