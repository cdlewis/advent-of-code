package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/cdlewis/advent-of-code/2024/util"
	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/cast"
	"github.com/cdlewis/advent-of-code/2024/util/grid"
)

var re = regexp.MustCompile("-?[0-9]+")

type Robot struct {
	Position grid.Point
	Velocity grid.Point
}

func main() {
	width := 101
	height := 103
	seconds := 10_000
	input := aoc.GetInput(14, false, "")

	var robots []Robot
	for _, l := range strings.Split(input, "\n") {
		n := re.FindAllString(l, -1)
		robots = append(robots, Robot{
			Position: grid.Point{cast.ToInt(n[1]), cast.ToInt(n[0])},
			Velocity: grid.Point{cast.ToInt(n[3]), cast.ToInt(n[2])},
		})
	}

	grid := make([][]int, height)
	for idx := range grid {
		grid[idx] = make([]int, width)
	}

	for _, r := range robots {
		grid[r.Position[0]][r.Position[1]]++
	}
	if len(robots) != 500 {
		panic("missing robot")
	}
	for s := range seconds {
		for idx := range robots {
			grid[robots[idx].Position[0]][robots[idx].Position[1]]--
			robots[idx].Position[0] = util.Mod(robots[idx].Position[0]+robots[idx].Velocity[0], height)
			robots[idx].Position[1] = util.Mod(robots[idx].Position[1]+robots[idx].Velocity[1], width)
			grid[robots[idx].Position[0]][robots[idx].Position[1]]++
		}

		fmt.Println("\033[2J")

		fmt.Println("Seconds = ", s)
		gridSum := 0
		for idx := range grid {
			for jdx := range grid[idx] {
				gridSum += grid[idx][jdx]
				if grid[idx][jdx] > 0 {
					fmt.Print(grid[idx][jdx])
				} else {
					fmt.Print("█")
				}
			}
			fmt.Println()
		}
		if gridSum != 500 {
			panic("Lost a robot!")
		}
		fmt.Println()
		fmt.Println()
		//fmt.Println(gridSum)
		fmt.Scanln()
	}
}

// 6492 too low
