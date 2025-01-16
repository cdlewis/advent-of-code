package eight

import (
	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/grid"
)

type Tile struct {
	start       grid.Point
	position    grid.Point
	direction   grid.Point
	antennaType byte
}

func Eight() int {
	uniquePoints := map[[2]int]struct{}{}
	grid := grid.ToByteGrid(aoc.GetInput(8, false, ""))
	for idx, i := range grid {
		for jdx, j := range i {
			if j == '.' {
				continue
			}

			explore([2]int{idx, jdx}, grid, uniquePoints)
		}
	}

	return len(uniquePoints)
}

func explore(start grid.Point, floorMap grid.Grid[byte], points map[[2]int]struct{}) {
	var q []Tile
	for _, d := range grid.DirectionsDiagonal {
		newPosition := start.Add(d)
		if !floorMap.ValidPoint(newPosition) {
			continue
		}

		q = append(q, Tile{
			start:       start,
			position:    newPosition,
			direction:   d,
			antennaType: floorMap.Get(start),
		})
	}

	seenSameAntenna := false
	seen := map[[2]int]struct{}{}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if _, ok := seen[curr.position]; ok {
			continue
		}
		seen[curr.position] = struct{}{}

		if floorMap.Get(curr.position) == curr.antennaType {
			seenSameAntenna = true

			moved := curr.position.Subtract(curr.start)

			firstPoint := curr.position.Add(moved)
			for floorMap.ValidPoint(firstPoint) {
				points[firstPoint] = struct{}{}
				firstPoint = firstPoint.Add(moved)
			}

			secondPoint := curr.start.Subtract(moved)
			for floorMap.ValidPoint(secondPoint) {
				points[secondPoint] = struct{}{}

				secondPoint = secondPoint.Subtract(moved)
			}
		}

		for _, d := range grid.DirectionsDiagonal {
			if curr.direction[0] != 0 && curr.direction[0] != d[0] {
				continue
			}

			if curr.direction[1] != 0 && curr.direction[1] != d[1] {
				continue
			}

			newPosition := curr.position.Add(d)
			if !floorMap.ValidPoint(newPosition) {
				continue
			}

			q = append(q, Tile{
				start:       curr.start,
				position:    newPosition,
				direction:   curr.direction,
				antennaType: curr.antennaType,
			})
		}
	}

	if seenSameAntenna {
		points[start] = struct{}{}
	}
}
