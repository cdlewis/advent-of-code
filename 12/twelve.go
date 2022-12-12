package twelve

import (
	"strings"

	"github.com/cdlewis/advent-of-code/util"
)

func Twelve(useTest bool, testInput string) int {
	graph := util.Map(strings.Split(util.GetInput(12, useTest, testInput), "\n"), func(s string) []byte { return []byte(s) })

	end := [2]int{}
	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] == 'E' {
				end = [2]int{i, j}
			}
		}
	}

	bestResult := -1

	shortestPath, exists := util.ShortestUnweightedPath(
		graph,
		end,
		func(x [2]int) bool { return graph[x[0]][x[1]] == 'a' || graph[x[0]][x[1]] == 'S' },
		func(x, y [2]int) bool { return normalize(graph[x[0]][x[1]]) <= normalize(graph[y[0]][y[1]])+1 },
	)

	if exists && (shortestPath < bestResult || bestResult == -1) {
		bestResult = shortestPath
	}

	return bestResult
}

func normalize(b byte) int {
	if b == 'S' {
		return int('a')
	}

	if b == 'E' {
		return int('z')
	}

	return int(b)
}
