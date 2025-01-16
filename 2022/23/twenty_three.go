package twenty_three

import (
	"math"
	"strings"

	"github.com/cdlewis/advent-of-code/2022/util"
)

func TwentyThree() int {
	raw := strings.Split(util.GetInput(23, false, `....#..
	..###.#
	#...#.#
	.#...##
	#.###..
	##.#.##
	.#..#..`), "\n")

	graph := map[[2]int]int{}
	directionsGroups := append([][][2]int{}, util.DirectionsDiagonalGrouped...)

	for idx := range raw {
		for jdx, j := range raw[idx] {
			if j == '#' {
				graph[[2]int{idx, jdx}] = -1
			}
		}
	}

	for i := 0; i < math.MaxInt; i++ {
		proposedMoves := map[[2]int][][2]int{}
		moveRequired := false

		for pos := range graph {
			if spaceEmpty(pos, util.DirectionsDiagonal, graph) {
				continue
			}

			moveRequired = true

			for _, group := range directionsGroups {
				if spaceEmpty(pos, group, graph) {
					newPos := util.AddPoints(pos, group[1])
					proposedMoves[newPos] = append(proposedMoves[newPos], pos)
					break
				}
			}
		}

		for newPos, proposedMovers := range proposedMoves {
			if len(proposedMovers) > 1 {
				continue
			}

			graph[newPos] = i

			existingPos := proposedMovers[0]
			if round := graph[existingPos]; round != i {
				delete(graph, existingPos)
			}
		}

		directionsGroups = util.RotateRight(directionsGroups)

		if !moveRequired {
			return i + 1
		}
	}

	panic("no solution")
}

func spaceEmpty(pos [2]int, directions [][2]int, graph map[[2]int]int) bool {
	empty := true

	for _, d := range directions {
		newPos := util.AddPoints(pos, d)
		if _, occupied := graph[newPos]; occupied {
			empty = false
			break
		}
	}

	return empty
}
