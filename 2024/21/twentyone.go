package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/grid"
)

var testInput = `029A
980A
179A
456A
379A`

var keypad = grid.Grid[byte]([][]byte{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{'#', '0', 'A'},
})

var directionPad = [][]byte{
	{'#', '^', 'A'},
	{'<', 'v', '>'},
}

var changeToDirection = map[grid.Point]byte{
	grid.UP:    '^',
	grid.DOWN:  'v',
	grid.LEFT:  '<',
	grid.RIGHT: '>',
}

func main() {
	input := strings.Split(aoc.GetInput(21, false, testInput), "\n")

	total := 0
	for _, row := range input {
		key, _ := strconv.Atoi(row[:len(row)-1])
		complexity := cost(row, -1)
		total += (key * complexity)
	}

	fmt.Println(total == 195664513288128)
}

var cache = map[string]int{}

func cost(code string, robot int) int {
	if result, ok := cache[code+strconv.Itoa(robot)]; ok {
		return result
	}

	maze := keypad
	if robot > 0 {
		maze = directionPad
	}

	totalCost := 0
	prev := 'A'
	for _, curr := range code {
		posA, _ := maze.Find(byte(prev))
		posB, _ := maze.Find(byte(curr))

		paths := allPaths(posA, posB, maze)
		cheapestPath := math.MaxInt
		for _, p := range paths {
			if robot == -1 {
				length := cost(p, 25)
				if length < cheapestPath {
					cheapestPath = length
				}
			} else if robot > 1 {
				length := cost(p, robot-1)
				if length < cheapestPath {
					cheapestPath = length
				}
			} else {
				if len(p) < cheapestPath {
					cheapestPath = len(p)
				}
			}
		}

		totalCost += cheapestPath
		prev = curr
	}

	cache[code+strconv.Itoa(robot)] = totalCost
	return totalCost
}

type item struct {
	path []grid.Point
	seen []grid.Point
}

func allPaths(
	a grid.Point,
	b grid.Point,
	maze grid.Grid[byte],
) []string {
	var results [][]grid.Point
	q := []item{{seen: []grid.Point{a}}}

	for len(q) > 0 {
		curr := q[0]
		currHead := curr.seen[len(curr.seen)-1]
		q = q[1:]

		if currHead == b {
			results = append(results, curr.path)
		}

		for _, d := range grid.Directions {
			nextPos := currHead.Add(d)
			if maze.GetOrElse(nextPos, '#') == '#' {
				continue
			}

			if slices.Index(curr.seen, nextPos) != -1 {
				continue
			}

			q = append(q, item{
				path: append(slices.Clone(curr.path), d),
				seen: append(slices.Clone(curr.seen), nextPos),
			})
		}
	}

	var adaptedResults []string
	for _, r := range results {
		var adaptedResult []byte
		for _, i := range r {
			adaptedResult = append(adaptedResult, changeToDirection[i])
		}
		adaptedResult = append(adaptedResult, 'A')

		adaptedResults = append(adaptedResults, string(adaptedResult))
	}

	return adaptedResults
}
