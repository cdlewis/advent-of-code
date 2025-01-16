package eighteen

import (
	"github.com/cdlewis/advent-of-code/2022/util"
)

const AIR = 0
const LAVA_EXPLORED = 1
const LAVA_UNEXPLORED = 2

func Eighteen(input string) int {
	rows := util.ToGrid(input)

	// Determine space dimension

	maxX, maxY, maxZ := 0, 0, 0
	for _, row := range rows {
		maxX = util.Max(maxX, row[0])
		maxY = util.Max(maxY, row[1])
		maxZ = util.Max(maxZ, row[2])
	}

	// Construct 3D space

	space := make([][][]int, maxX+1)
	for i := range space {
		space[i] = make([][]int, maxY+1)
		for j := range space[i] {
			space[i][j] = make([]int, maxZ+1)
		}
	}

	// Mark lava droplets

	for _, r := range rows {
		space[r[0]][r[1]][r[2]] = LAVA_UNEXPLORED
	}

	// Fill in any air bubbles

	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			for k := 0; k < maxZ; k++ {
				// Very inefficient but it works
				seen := map[[3]int]bool{}

				if space[i][j][k] != AIR {
					continue
				}

				q := [][3]int{{i, j, k}}

				foundEdge := false
			BFS:
				for len(q) > 0 {
					curr := q[0]
					q = q[1:]

					if seen[curr] {
						continue
					}
					seen[curr] = true

					for _, d := range util.Directions3D {
						nX := curr[0] + d[0]
						nY := curr[1] + d[1]
						nZ := curr[2] + d[2]

						if !util.ValidCoordinate3D(nX, nY, nZ, space) {
							foundEdge = true
							break BFS
						}

						if space[nX][nY][nZ] == AIR {
							q = append(q, [3]int{nX, nY, nZ})
						}
					}
				}

				// Air hole detected!
				if !foundEdge {
					for coord := range seen {
						space[coord[0]][coord[1]][coord[2]] = 2
					}
				}
			}
		}
	}

	surfaces := 0

	for _, r := range rows {
		q := [][]int{r}

		for len(q) > 0 {
			curr := q[0]
			q = q[1:]

			if space[curr[0]][curr[1]][curr[2]] <= 1 {
				continue
			}

			space[curr[0]][curr[1]][curr[2]] = 1

			for _, d := range util.Directions3D {
				nX := curr[0] + d[0]
				nY := curr[1] + d[1]
				nZ := curr[2] + d[2]

				if !util.ValidCoordinate3D(nX, nY, nZ, space) || space[nX][nY][nZ] == 0 {
					surfaces++
				} else {
					q = append(q, []int{nX, nY, nZ})
				}
			}
		}
	}

	return surfaces
}
