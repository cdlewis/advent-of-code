package grid

import (
	"fmt"
	"math"
	"strings"

	"github.com/cdlewis/advent-of-code/2024/util/cast"
)

type Point [2]int

func (p Point) Add(another Point) Point {
	return AddPoints(p, another)
}

func (p Point) Subtract(another Point) Point {
	return SubtractPoints(p, another)
}

func (p Point) ToDirection() []Point {
	var results []Point

	if p[0] > 0 {
		results = append(results, DOWN)
	} else if p[0] < 0 {
		results = append(results, UP)
	}

	if p[1] > 0 {
		results = append(results, RIGHT)
	} else if p[1] < 0 {
		results = append(results, LEFT)
	}

	return results
}

func (p Point) RotateClockwise() Point {
	switch p {
	case LEFT:
		return UP
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	}

	panic("impossible state")
}

func (p Point) RotateCounterClockwise() Point {
	switch p {
	case LEFT:
		return DOWN
	case DOWN:
		return RIGHT
	case RIGHT:
		return UP
	case UP:
		return LEFT
	}

	panic("impossible state")
}

type Grid[T comparable] [][]T

func (g Grid[T]) ValidPoint(point Point) bool {
	return ValidPointCoordinate(point, g)
}

func (g Grid[T]) Get(point Point) T {
	return g[point[0]][point[1]]
}

func (g Grid[T]) GetOrElse(point Point, orElse T) T {
	if !g.ValidPoint(point) {
		return orElse
	}

	return g.Get(point)
}

func (g Grid[T]) Set(point Point, val T) {
	g[point[0]][point[1]] = val
}

func (g Grid[T]) GetAdjacent(point Point) []Point {
	var result []Point
	for _, i := range Directions {
		newPoint := point.Add(i)
		if g.ValidPoint(newPoint) {
			result = append(result, newPoint)
		}
	}
	return result
}

func (g Grid[T]) Find(needle T) (Point, bool) {
	for idx := range g {
		for jdx := range g[idx] {
			if g[idx][jdx] == needle {
				return Point{idx, jdx}, true
			}
		}
	}

	return Point{}, false
}

func (g Grid[T]) Print() {
	for _, i := range g {
		for _, j := range i {
			byteVal, ok := any(j).(byte)
			if ok {
				fmt.Print(string(byteVal))
			} else {
				fmt.Print(j)
			}
		}
		fmt.Println()
	}
}

func ValidCoordinate[U any](i int, j int, grid [][]U) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}

func ValidPointCoordinate[U any](point [2]int, grid [][]U) bool {
	return ValidCoordinate(point[0], point[1], grid)
}

var UP = Point{-1, 0}
var DOWN = Point{1, 0}
var LEFT = Point{0, -1}
var RIGHT = Point{0, 1}

var Directions = []Point{UP, DOWN, LEFT, RIGHT}

var DirectionsDiagonal = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
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

func ShortestUnweightedPath[U comparable](
	graph Grid[U],
	start Point,
	isEnd func(Point) bool,
	validatePath func(Point, Point) bool,
) (int, bool) {
	steps := 0
	stack := []Point{start}
	visited := map[Point]bool{}

	for len(stack) > 0 {
		var newStack []Point

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
				nextCoord := curr.Add(d)
				if graph.ValidPoint(nextCoord) && validatePath(curr, nextCoord) {
					newStack = append(newStack, nextCoord)
				}
			}
		}

		stack = newStack
		steps++
	}

	return -1, false
}

func ToGrid(s string) Grid[int] {
	lines := strings.Split(s, "\n")
	result := make([][]int, 0, len(lines))

	for _, l := range lines {
		line := make([]int, 0, len(l))
		for _, j := range l {
			line = append(line, cast.ToInt(j))
		}
		result = append(result, line)
	}

	return result
}

func ToByteGrid(s string) Grid[byte] {
	lines := strings.Split(s, "\n")
	result := make([][]byte, 0, len(lines))

	for _, l := range lines {
		result = append(result, []byte(l))
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
	return [2]int{x[0] + y[0], x[1] + y[1]}
}

func SubtractPoints(x, y [2]int) [2]int {
	return [2]int{x[0] - y[0], x[1] - y[1]}
}

func BoundingBox[U any](graph map[[2]int]U) (int, int, int, int) {
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for pos := range graph {
		minY = min(minY, pos[0])
		maxY = max(maxY, pos[0])
		minX = min(minX, pos[1])
		maxX = max(maxX, pos[1])
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
