package util

import (
	"regexp"
	"strings"
)

func ValidCoordinate[U any](i int, j int, grid [][]U) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}

var Directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func ShortestUnweightedPath[U any](graph [][]U, start [2]int, isEnd func(x [2]int) bool, validatePath func(x [2]int, y [2]int) bool) (int, bool) {
	steps := 0
	stack := [][2]int{start}
	visited := map[[2]int]bool{}

	for len(stack) > 0 {
		newStack := [][2]int{}

		for len(stack) > 0 {
			curr := stack[0]
			stack = stack[1:]

			if isEnd(curr) {
				return steps, true
			}

			if visited[curr] {
				continue
			}

			visited[curr] = true

			for _, d := range Directions {
				nextCoord := [2]int{curr[0] + d[0], curr[1] + d[1]}
				if ValidCoordinate(curr[0]+d[0], curr[1]+d[1], graph) && validatePath(curr, nextCoord) {
					newStack = append(newStack, nextCoord)
				}
			}
		}

		stack = newStack
		steps++
	}

	return -1, false
}

func ToGrid(s string) [][]int {
	lines := strings.Split(s, "\n")
	result := make([][]int, len(lines))

	re := regexp.MustCompile(`(-?[0-9])+`)

	for idx, l := range lines {
		result[idx] = Map(re.FindAllString(l, -1), ToInt[string])
	}

	return result
}

var Directions3D = [][]int{
	{1, 0, 0},
	{-1, 0, 0},
	{0, 1, 0},
	{0, -1, 0},
	{0, 0, 1},
	{0, 0, -1},
}

func ValidCoordinate3D[U any](i, j, k int, space [][][]U) bool {
	if i < 0 || i >= len(space) {
		return false
	}

	if j < 0 || j >= len(space[i]) {
		return false
	}

	if k < 0 || k >= len(space[i][j]) {
		return false
	}

	return true
}
