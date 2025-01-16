package main

import (
	"fmt"

	"github.com/cdlewis/advent-of-code/2024/util"
	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/grid"
)

func main() {
	input := grid.ToByteGrid(aoc.GetInput(20, false, ""))

	var start, end grid.Point
	for idx, i := range input {
		for jdx, j := range i {
			if j == 'S' {
				start = grid.Point{idx, jdx}
				input[idx][jdx] = '.'
			}

			if j == 'E' {
				end = grid.Point{idx, jdx}
				input[idx][jdx] = '.'
			}
		}
	}

	cheatLength := 20
	minSavings := 100

	shortestWithoutCheating, _ := grid.ShortestUnweightedPath(
		input,
		start,
		func(p grid.Point) bool { return p == end },
		func(p1, p2 grid.Point) bool { return input.Get(p2) != '#' },
	)

	toStart := calculateDistsFrom(start, input)
	toEnd := calculateDistsFrom(end, input)

	results := 0
	for p1, p1Dist := range toStart {
		for p2, p2Dist := range toEnd {
			p1p2Dist := util.Abs(p1[0]-p2[0]) + util.Abs(p1[1]-p2[1])
			if p1p2Dist > cheatLength {
				continue
			}

			totalDist := p1Dist + p1p2Dist + p2Dist

			if shortestWithoutCheating-totalDist >= minSavings {
				results++
			}
		}
	}

	fmt.Println(results)
}

func calculateDistsFrom(start grid.Point, maze grid.Grid[byte]) map[grid.Point]int {
	q := []grid.Point{start}
	dist := 0
	result := map[grid.Point]int{}

	for len(q) > 0 {
		var newQ []grid.Point

		for len(q) > 0 {
			curr := q[0]
			q = q[1:]

			if _, ok := result[curr]; ok {
				continue
			}
			result[curr] = dist

			for _, newPos := range maze.GetAdjacent(curr) {
				if maze.Get(newPos) != '#' {
					newQ = append(newQ, newPos)
				}
			}
		}
		dist++
		q = newQ
	}

	return result
}
