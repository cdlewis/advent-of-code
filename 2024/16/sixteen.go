package main

import (
	"fmt"
	"math"

	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/grid"
	"github.com/cdlewis/advent-of-code/2024/util/set"
)

func main() {
	maze := grid.ToByteGrid(aoc.GetInput(16, false, ""))

	currentDirection := grid.RIGHT
	var start, end grid.Point
	for idx := range maze {
		for jdx := range maze[idx] {
			if maze[idx][jdx] == 'S' {
				start = grid.Point{idx, jdx}
			}

			if maze[idx][jdx] == 'E' {
				end = grid.Point{idx, jdx}
			}
		}
	}

	_, pointSet, found := shortestPath(start, end, currentDirection, maze, 0)
	if !found {
		panic("no path found")
	}

	fmt.Println(len(pointSet))
}

type Choice struct {
	Point     grid.Point
	Cost      int
	Direction grid.Point
}

var seen = map[[4]int]int{}
var bestDiscoveredCost = math.MaxInt

func shortestPath(
	from grid.Point,
	to grid.Point,
	direction grid.Point,
	maze grid.Grid[byte],
	currentCost int,
) (int, set.Set[grid.Point], bool) {
	if currentCost > bestDiscoveredCost {
		return -1, nil, false
	}

	key := [4]int{from[0], from[1], direction[0], direction[1]}
	if cost, found := seen[key]; found && cost < currentCost {
		return -1, nil, false
	}
	seen[key] = currentCost

	if from == to {
		bestDiscoveredCost = min(bestDiscoveredCost, currentCost)
		return currentCost, set.New(from), true
	}

	var choices []Choice
	for _, d := range grid.Directions {
		nextPoint := from.Add(d)

		if d == direction {
			choices = append(choices, Choice{
				Point:     nextPoint,
				Cost:      currentCost + 1,
				Direction: direction,
			})
			continue
		}

		// Don't bother facing a wall and trying to explore
		if maze.Get(nextPoint) == '#' {
			continue
		}

		cost := currentCost + 1001
		if d == direction.RotateClockwise().RotateClockwise() {
			cost += 1000
		}

		// We already found something cheaper
		if cost > bestDiscoveredCost {
			continue
		}

		choices = append(choices, Choice{
			Point:     nextPoint,
			Cost:      cost,
			Direction: d,
		})
	}

	var finalCosts []int
	var finalPoints []set.Set[grid.Point]
	for _, c := range choices {
		if maze.ValidPoint(c.Point) && maze.Get(c.Point) != '#' && c.Cost <= bestDiscoveredCost {
			cost, points, found := shortestPath(c.Point, to, c.Direction, maze, c.Cost)
			if found {
				finalCosts = append(finalCosts, cost)
				finalPoints = append(finalPoints, points)
			}
		}
	}

	if len(finalCosts) == 0 {
		return -1, nil, false
	}

	minIdxs := []int{0}
	for idx, i := range finalCosts {
		if i < finalCosts[minIdxs[0]] {
			minIdxs = []int{idx}
		} else if i == finalCosts[minIdxs[0]] {
			minIdxs = append(minIdxs, idx)
		}
	}

	pointSet := set.New(from)
	for _, i := range minIdxs {
		pointSet.Extend(finalPoints[i])
	}

	return finalCosts[minIdxs[0]], pointSet, true
}
