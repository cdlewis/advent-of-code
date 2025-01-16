package nine

import (
	"fmt"
	"strings"

	"github.com/cdlewis/advent-of-code/2022/util"
)

var upDownLeftRight = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var diagonal = [][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func Nine(useTest bool, test string) int {
	instructions := strings.Split(util.GetInput(9, useTest, test), "\n")

	knots := make([][2]int, 10)
	visited := map[[2]int]bool{{0, 0}: true}

	for _, instruction := range instructions {
		dx, dy := getDirection(instruction[0])
		distance := util.ToInt(instruction[2:])

		for i := 0; i < distance; i++ {
			// Store the current tail entry
			visited[knots[len(knots)-1]] = true

			// Update head of the rope
			knots[0][0] += dx
			knots[0][1] += dy

			for i := 1; i < len(knots); i++ {
				previousX, previousY := knots[i-1][0], knots[i-1][1]
				currentX, currentY := knots[i][0], knots[i][1]

				if util.Abs(previousX-currentX) <= 1 && util.Abs(previousY-currentY) <= 1 {
					continue
				}

				// Determine which directions to search for the next position
				possibleMovements := upDownLeftRight
				if previousX != currentX && previousY != currentY {
					possibleMovements = diagonal
				}

				for _, p := range possibleMovements {
					nextX := currentX + p[0]
					nextY := currentY + p[1]

					if util.Abs(previousX-nextX) <= 1 && util.Abs(previousY-nextY) <= 1 {
						knots[i][0] = nextX
						knots[i][1] = nextY
						break
					}
				}
			}
		}
	}
	fmt.Println(len(visited))
	return len(visited)
}

func getDirection(d byte) (int, int) {
	switch d {
	case 'U':
		return 0, 1
	case 'D':
		return 0, -1
	case 'L':
		return -1, 0
	case 'R':
		return 1, 0
	default:
		panic("Unexpected input")
	}
}
