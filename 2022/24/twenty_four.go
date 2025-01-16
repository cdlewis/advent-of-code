package twenty_four

import (
	"math"
	"strings"

	"github.com/cdlewis/advent-of-code/2022/util"
)

func TwentyFour() int {
	maze := parseInput()

	// Model the blizzards over time (there must be a cycle) and build an adjacency
	// map describing *when* you can safely move to a particular point
	safe, maxCycle := adjacencyMapFromBlizzards(maze)

	startToGoal := journey(maze.Start, maze.End, 0, maxCycle, safe, maze)
	backToStart := journey(maze.End, maze.Start, startToGoal, maxCycle, safe, maze)
	backToGoal := journey(maze.Start, maze.End, backToStart, maxCycle, safe, maze)

	return backToGoal
}

type Maze struct {
	Blizzards []Blizzard
	Start     [2]int
	End       [2]int
	MinX      int
	MaxX      int
	MinY      int
	MaxY      int
}

type Blizzard struct {
	Pos       [2]int
	Direction [2]int
}

func parseInput() Maze {
	raw := strings.Split(util.GetInput(24, false, ``), "\n")
	blizzards := []Blizzard{}
	maxY, maxX := len(raw)-1, len(raw[0])-1
	start, end := [2]int{}, [2]int{}

	for idx := range raw {
		for jdx, j := range raw[idx] {
			switch j {
			case '>':
				blizzards = append(blizzards, Blizzard{
					Pos:       [2]int{idx, jdx},
					Direction: [2]int{0, 1},
				})
			case '<':
				blizzards = append(blizzards, Blizzard{
					Pos:       [2]int{idx, jdx},
					Direction: [2]int{0, -1},
				})
			case '^':
				blizzards = append(blizzards, Blizzard{
					Pos:       [2]int{idx, jdx},
					Direction: [2]int{-1, 0},
				})
			case 'v':
				blizzards = append(blizzards, Blizzard{
					Pos:       [2]int{idx, jdx},
					Direction: [2]int{1, 0},
				})
			case '.':
				if idx == len(raw)-1 {
					end = [2]int{idx, jdx}
				} else if idx == 0 {
					start = [2]int{idx, jdx}
				}
			}
		}
	}

	return Maze{
		MinX:      1,
		MinY:      1,
		MaxX:      maxX,
		MaxY:      maxY,
		Blizzards: blizzards,
		Start:     start,
		End:       end,
	}
}

func adjacencyMapFromBlizzards(maze Maze) (map[[3]int]bool, int) {
	notSafeWhen := map[[2]int][]int{}
	seenC := map[string]bool{}
	maxCycle := 0

	for i := 0; ; i++ {
		cycleReached := false
		builder := strings.Builder{}
		builder.Grow(maze.MaxY * maze.MaxX)

		for y := 1; y < maze.MaxY; y++ {
			for x := 1; x < maze.MaxX; x++ {
				draw := 0

				for _, b := range maze.Blizzards {
					if b.Pos[0] == y && b.Pos[1] == x {
						draw++
					}
				}

				if draw == 0 {
					builder.WriteString(".")
				} else {
					builder.WriteString(string(rune(draw)))
				}
			}
		}

		key := builder.String()
		if seenC[key] {
			cycleReached = true
		}
		seenC[key] = true

		for idx, b := range maze.Blizzards {
			notSafeWhen[b.Pos] = append(notSafeWhen[b.Pos], i)

			maze.Blizzards[idx].Pos[0] = b.Pos[0] + b.Direction[0]
			if maze.Blizzards[idx].Pos[0] == 0 {
				maze.Blizzards[idx].Pos[0] = maze.MaxY - 1
			}
			if maze.Blizzards[idx].Pos[0] == maze.MaxY {
				maze.Blizzards[idx].Pos[0] = 1
			}
			maze.Blizzards[idx].Pos[1] = b.Pos[1] + b.Direction[1]
			if maze.Blizzards[idx].Pos[1] == 0 {
				maze.Blizzards[idx].Pos[1] = maze.MaxX - 1
			}
			if maze.Blizzards[idx].Pos[1] == maze.MaxX {
				maze.Blizzards[idx].Pos[1] = 1
			}
		}

		if cycleReached {
			maxCycle = i
			break
		}
	}

	safe := map[[3]int]bool{}
	for i := 1; i < maze.MaxY; i++ {
		for j := 1; j < maze.MaxX; j++ {
			for k := 0; k <= maxCycle; k++ {
				pos := [2]int{i, j}
				time := k
				posTime := [3]int{i, j, k}

				safeNow := true
				for _, l := range notSafeWhen[pos] {
					if l == time {
						safeNow = false
						break
					}
				}

				if safeNow {
					safe[posTime] = true
				}
			}
		}
	}

	return safe, maxCycle
}

func journey(start [2]int, target [2]int, currentCycle, maxCycle int, safe map[[3]int]bool, maze Maze) int {
	best := math.MaxInt
	s := [][3]int{{start[0], start[1], currentCycle}}
	seen := map[[3]int]int{}

	for len(s) > 0 {
		curr := s[0]
		s = s[1:]

		if curr[2] > best {
			continue
		}

		key := curr
		key[2] = key[2] % maxCycle
		if when, ok := seen[key]; ok && curr[2] >= when {
			continue
		}
		seen[key] = curr[2]

		for _, d := range util.Directions {
			newY := curr[0] + d[0]
			newX := curr[1] + d[1]

			if newY == target[0] && newX == target[1] {
				best = util.Min(curr[2]+1, best)
				continue
			}

			if newX <= 0 || newX >= maze.MaxX || newY <= 0 || newY >= maze.MaxY || (newY == start[0] && newX == start[1]) {
				continue
			}

			for i := 1; i < maxCycle; i++ {
				newPos := [3]int{newY, newX, curr[2] + i}
				newPosKey := [3]int{newY, newX, util.Mod(curr[2]+i, maxCycle)}
				if safe[newPosKey] {
					s = append(s, newPos)
				} else {
					if curr[0] != start[0] && curr[1] != start[1] {
						break // can't hold pos here
					}
				}
			}
		}
	}

	return best
}
