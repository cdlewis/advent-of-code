package twelve

import (
	"strings"

	"github.com/cdlewis/advent-of-code/util"
)

func Twelve(useTest bool, testInput string) int {
	graph := util.Map(strings.Split(util.GetInput(12, useTest, testInput), "\n"), func(s string) []byte { return []byte(s) })

	end := [2]int{}
	for idx := range graph {
		for jdx, cell := range graph[i] {
			if cell == 'E' {
				end = [2]int{idx, jdx}
			}
		}
	}

	shortestPath, _ := util.ShortestUnweightedPath(
		graph,
		end,
		func(x [2]int) bool { return graph[x[0]][x[1]] == 'a' || graph[x[0]][x[1]] == 'S' },
		func(x, y [2]int) bool { return normalize(graph[x[0]][x[1]]) <= normalize(graph[y[0]][y[1]])+1 },
	)

	return shortestPath
}

func normalize(b byte) int {
	switch b {
	case 'S':
		return int('a')
	case 'E':
		return int('z')
	default:
		return int(b)
	}
}
