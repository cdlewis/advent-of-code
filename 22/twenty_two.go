package twenty_two

import (
	"strings"

	"github.com/cdlewis/advent-of-code/util"
)

func TwentyTwo() int {
	surface, instructions := parseInput()

	x, y := 0, 0
	for i := 0; i < len(surface[0]); i++ {
		if surface[y][i] == '.' {
			x = i
			break
		}
	}

	directionIndex := RightIndex

	for _, i := range instructions {
		switch i {
		case "L":
			directionIndex = util.Mod(directionIndex-1, len(directionClock))
		case "R":
			directionIndex = util.Mod(directionIndex+1, len(directionClock))
		default:
			move := util.ToInt(i)

			for j := 0; j < move; j++ {
				dY := util.Mod(y+directionClock[directionIndex][0], len(surface))
				dX := util.Mod(x+directionClock[directionIndex][1], len(surface[x]))

				// Out of bounds -- try to find a place to land
				if surface[dY][dX] == ' ' {
					newdY, newdX, newDirectionIndex := translateCoordinate(y, x, directionIndex, 50)

					// Respect the wall
					if surface[newdY][newdX] == '#' {
						break
					}

					y, x = newdY, newdX
					directionIndex = newDirectionIndex
					continue
				}

				// wall -- stop moving
				if surface[dY][dX] == '#' {
					break
				}

				// record new position
				y, x = dY, dX
				x = dX
			}

		}
	}

	return 1000*(y+1) + 4*(x+1) + directionIndex
}

var directionClock = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

const RightIndex = 0
const DownIndex = 1
const LeftIndex = 2
const UpIndex = 3

func parseInput() ([]string, []string) {
	raw := strings.Split(util.GetInput(22, false, ``), "\n")

	rawinstructions := raw[len(raw)-1]
	instructions := []string{}
	i := 0
	for i < len(rawinstructions) {
		if rawinstructions[i] == 'L' || rawinstructions[i] == 'R' {
			instructions = append(instructions, string(rawinstructions[i]))
			i++
			continue
		}

		j := i
		for j < len(rawinstructions) && rawinstructions[j] != 'L' && rawinstructions[j] != 'R' {
			j++
		}

		instructions = append(instructions, rawinstructions[i:j])
		i = j
	}

	raw = raw[:len(raw)-2]

	maxLength := 0
	for _, i := range raw {
		if len(i) > maxLength {
			maxLength = len(i)
		}
	}

	for idx := range raw {
		if len(raw[idx]) < maxLength {
			raw[idx] = util.PadRight(raw[idx], maxLength, ' ')
		}
	}

	return raw, instructions
}

// fml
func translateCoordinate(y, x int, direction int, gridSize int) (int, int, int) {
	yOffset, xOffset := util.Mod(y, 50), util.Mod(x, 50)
	yGridPosition, xGridPosition := y/50, x/50

	switch direction {
	case RightIndex:
		if yGridPosition == 0 && xGridPosition == 2 {
			return 3*gridSize - 1 - yOffset, 2*gridSize - 1, LeftIndex
		} else if yGridPosition == 2 && xGridPosition == 1 {
			return gridSize - 1 - yOffset, 3*gridSize - 1, LeftIndex
		} else if yGridPosition == 1 && xGridPosition == 1 {
			return gridSize - 1, 2*gridSize + yOffset, UpIndex
		} else if yGridPosition == 3 && xGridPosition == 0 {
			return 3*gridSize - 1, gridSize + yOffset, UpIndex
		}

	case DownIndex:
		if yGridPosition == 3 && xGridPosition == 0 {
			return 0, 2*gridSize + xOffset, DownIndex
		} else if yGridPosition == 0 && xGridPosition == 2 {
			return gridSize + xOffset, 2*gridSize - 1, LeftIndex
		} else if yGridPosition == 2 && xGridPosition == 1 {
			return 3*gridSize + xOffset, gridSize - 1, LeftIndex
		}

	case LeftIndex:
		if yGridPosition == 3 && xGridPosition == 0 {
			return 0, gridSize + yOffset, DownIndex
		} else if yGridPosition == 0 && xGridPosition == 1 {
			return 3*gridSize - 1 - yOffset, 0, RightIndex
		} else if yGridPosition == 2 && xGridPosition == 0 {
			return gridSize - 1 - yOffset, gridSize, RightIndex
		} else if yGridPosition == 1 && xGridPosition == 1 {
			return 2 * gridSize, yOffset, DownIndex
		}

	case UpIndex:
		if yGridPosition == 0 && xGridPosition == 1 {
			return 3*gridSize + xOffset, 0, RightIndex
		} else if yGridPosition == 2 && xGridPosition == 0 {
			return gridSize + xOffset, gridSize, RightIndex
		} else if yGridPosition == 0 && xGridPosition == 2 {
			return 4*gridSize - 1, xOffset, UpIndex
		}
	}

	panic("unexpected output?")
}
