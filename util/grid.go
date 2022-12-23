package util

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func ValidCoordinate[U any](i int, j int, grid [][]U) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}

var Directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

var DirectionsDiagonal = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{1, -1}, {1, 0}, {1, 1},
	{1, -1}, {0, -1}, {-1, -1},
	{-1, 1}, {0, 1}, {-1, 1},
}

var DirectionsDiagonalGrouped = [][][2]int{
	{{-1, -1}, {-1, 0}, {-1, 1}},
	{{1, -1}, {1, 0}, {1, 1}},
	{{1, -1}, {0, -1}, {-1, -1}},
	{{-1, 1}, {0, 1}, {1, 1}},
}

var Directions3D = [][]int{
	{1, 0, 0},
	{-1, 0, 0},
	{0, 1, 0},
	{0, -1, 0},
	{0, 0, 1},
	{0, 0, -1},
}

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

func AddPoints(x, y [2]int) [2]int {
	x[0] += y[0]
	x[1] += y[1]
	return x
}

func BoundingBox[U any](graph map[[2]int]U) (int, int, int, int) {
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for pos := range graph {
		minY = Min(minY, pos[0])
		maxY = Max(maxY, pos[0])
		minX = Min(minX, pos[1])
		maxX = Max(maxX, pos[1])
	}
	return minX, minY, maxX, maxY
}

func Print[U any](graph map[[2]int]U) {
	minX, minY, maxX, maxY := BoundingBox(graph)

	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			if _, ok := graph[[2]int{i, j}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}
