package twelve

import (
	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/grid"
	"github.com/cdlewis/advent-of-code/2024/util/set"
)

func Twelve() int {
	garden := grid.ToByteGrid(aoc.GetInput(12, false, ""))

	seen := set.New[grid.Point]()

	totalCost := 0

	for idx, i := range garden {
		for jdx := range i {
			curr := grid.Point{idx, jdx}
			if seen.Exists(curr) {
				continue
			}

			totalCost += getCost(curr, garden, seen)
		}
	}

	return totalCost
}

func getCost(start grid.Point, garden grid.Grid[byte], seen set.Set[grid.Point]) int {
	plantType := garden.Get(start)
	area := 0

	faces := map[[4]int][]set.Set[grid.Point]{}

	q := []grid.Point{start}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if seen.Exists(curr) {
			continue
		}
		seen.Add(curr)

		area++

		for _, d := range grid.Directions {
			newPlot := curr.Add(d)

			if garden.ValidPoint(newPlot) && garden.Get(newPlot) == plantType {
				q = append(q, newPlot)
				continue
			}

			newFace := [4]int{d[0], d[1], curr[0], newPlot[0]}
			if d[0] == 0 {
				newFace = [4]int{d[0], d[1], curr[1], newPlot[1]}
			}

			matches := []int{}
			for idx, i := range faces[newFace] {
				for _, a := range garden.GetAdjacent(curr) {
					if i.Exists(a) {
						matches = append(matches, idx)
					}
				}
			}

			if len(matches) == 0 {
				faces[newFace] = append(faces[newFace], set.New(curr))
				continue
			}

			// Edge case: we have a new point connecting two or more existing match sets
			var newPointsSlice []set.Set[grid.Point]
			pointsToMerge := set.New(matches...)
			mergedSet := set.New(curr)
			for idx, i := range faces[newFace] {
				if pointsToMerge.Exists(idx) {
					mergedSet = mergedSet.Combine(i)
					continue
				}

				newPointsSlice = append(newPointsSlice, i)
			}
			newPointsSlice = append(newPointsSlice, mergedSet)
			faces[newFace] = newPointsSlice
		}
	}

	sides := 0
	for _, face := range faces {
		sides += len(face)
	}

	return area * sides
}
