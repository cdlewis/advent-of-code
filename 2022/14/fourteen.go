package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/cdlewis/advent-of-code/2022/util"
)

var directions = [][2]int{{0, 1}, {-1, 1}, {1, 1}}

func main() {
	raw := strings.Split(util.GetInput(14, false, `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`), "\n")

	// Parse all coord pairs and extract largest Y
	coords := [][][2]int{}
	largestY := -1
	for _, i := range raw {
		pairs := strings.Split(i, " -> ")
		coordSet := [][2]int{}

		for _, j := range pairs {
			p := strings.Split(j, ",")

			coordSet = append(coordSet, [2]int{util.ToInt(p[0]), util.ToInt(p[1])})
			largestY = util.Max(util.ToInt(p[1]), largestY)
		}

		coords = append(coords, coordSet)
	}
	largestY += 2 // bump by 2 per problem requirements

	// Build a matrix based on the largest expected values
	c := make([][]int, largestY+1)
	for i := range c {
		c[i] = make([]int, 800)
	}

	// Introduce a 'floor' for the room
	for i := 0; i < 800; i++ {
		c[largestY][i] = 1
	}

	// Populate the room based on the coords
	for _, coordSet := range coords {
		for idx, curr := range coordSet {
			if idx == 0 {
				continue
			}

			prev := coordSet[idx-1]

			// One of these loops will immediately terminate based on the direction indicated by the coords

			startX, endX := util.Min(prev[0], curr[0]), util.Max(prev[0], curr[0])
			for x := startX; x <= endX; x++ {
				c[prev[1]][x] = 1
			}

			startY, endY := util.Min(prev[1], curr[1]), util.Max(prev[1], curr[1])
			for y := startY; y <= endY; y++ {
				c[y][prev[0]] = 1
			}
		}
	}

	sand := 0

	// Clear the screen and create some room
	fmt.Print("\033[2J")

	for {
		printWorld(sand, c)
		time.Sleep(10 * time.Millisecond)

		sand++
		if c[0][500] == 2 {
			return
		}

		sandPos := []int{0, 500}
		c[0][500] = 2

	MovingSand:
		for {
			for _, d := range directions {
				newX, newY := sandPos[1]+d[0], sandPos[0]+d[1]
				if util.ValidCoordinate(newY, newX, c) && c[newY][newX] == 0 {
					c[sandPos[0]][sandPos[1]] = 0 // Important that we clear the cell which used to contain sand
					sandPos = []int{newY, newX}
					c[sandPos[0]][sandPos[1]] = 2

					continue MovingSand
				}
			}

			// Stop moving sand if no direction is valid
			break
		}
	}
}

func printWorld(round int, c [][]int) {
	fmt.Printf("\033[%d;%dH", 0, 0)
	fmt.Println("Round: ", round)

	for _, i := range c {
		for _, j := range i[400:572] {
			if j == 0 {
				fmt.Print(" ")
			} else if j == 1 {
				fmt.Print("🪨")
			} else {
				fmt.Print("❄️")
			}
		}
		fmt.Print("\n")
	}
}
