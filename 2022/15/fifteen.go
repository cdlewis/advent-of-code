package fifteen

import (
	"github.com/cdlewis/advent-of-code/2022/util"
)

func Fifteen() int {
	grid := util.ToGrid(util.GetInput(15, false, ``))

	// Pre-compute some helpful things
	beaconLocations := map[[2]int]bool{}
	cachedManhattanDistances := make([]int, len(grid))
	for idx, i := range grid {
		beaconLocations[[2]int{i[2], i[3]}] = true
		cachedManhattanDistances[idx] = manhattan(i[0], i[1], i[2], i[3])
	}

	min, max := 0, 4_000_000
	for y := min; y < max; y++ {
	ScanXCoord:
		for x := min; x <= max; x++ {
			if _, ok := beaconLocations[[2]int{x, y}]; ok {
				continue
			}

			for idx, i := range grid {
				distanceToSensor := manhattan(i[0], i[1], x, y)

				if distance := distanceToSensor - cachedManhattanDistances[idx]; distance <= 0 {
					if jumpAhead := util.Abs(distance); jumpAhead > 1 {
						x += (jumpAhead - 1)
					}

					continue ScanXCoord
				}
			}

			return x*4000000 + y
		}
	}

	panic("No solution")
}

func manhattan(x1, y1, x2, y2 int) int {
	return util.Abs(x1-x2) + util.Abs(y1-y2)
}
