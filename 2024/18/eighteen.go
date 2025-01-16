package main

import (
	"fmt"
	"strings"

	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/cast"
	"github.com/cdlewis/advent-of-code/2024/util/grid"
)

var testInput = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

const testSize = 71

const numBytes = 1024

func main() {
	start := grid.Point{0, 0}
	end := grid.Point{testSize - 1, testSize - 1}
	raw := strings.Split(aoc.GetInput(18, false, testInput), "\n")

	world := make(grid.Grid[byte], testSize)
	for i := range world {
		world[i] = make([]byte, testSize)
	}

	for _, i := range raw {
		tokens := cast.FindAllInt(i)
		world[tokens[1]][tokens[0]] = '#'

		_, found := grid.ShortestUnweightedPath(
			world,
			start,
			func(p grid.Point) bool { return p == end },
			func(a, b grid.Point) bool { return world.Get(b) != '#' },
		)

		if !found {
			fmt.Println(i)
			break
		}
	}
}
