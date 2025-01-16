package main

import (
	"fmt"
	"strings"
	"time"

	gocursor "github.com/ahmetb/go-cursor"
	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/grid"
)

var replacer = strings.NewReplacer(
	"#", "##",
	"O", "[]",
	".", "..",
	"@", "@.",
)

var instructionMap = map[rune]grid.Point{
	'<': grid.LEFT,
	'^': grid.UP,
	'v': grid.DOWN,
	'>': grid.RIGHT,
}

func main() {
	rawInput := strings.Split(aoc.GetInput(15, false, ""), "\n\n")
	warehouse := grid.ToByteGrid(replacer.Replace(rawInput[0]))
	instructions := rawInput[1]

	var currentPosition grid.Point
START_SEARCH:
	for idx, i := range warehouse {
		for jdx := range i {
			if warehouse[idx][jdx] == '@' {
				currentPosition = grid.Point{idx, jdx}
				break START_SEARCH
			}
		}
	}

	for _, i := range instructions {
		time.Sleep(30 * time.Millisecond)
		dump(warehouse)
		gocursor.MoveUp(len(warehouse))
		gocursor.MoveLeft(len(warehouse[0]))
		dP := instructionMap[i]
		newPosition := currentPosition.Add(dP)

		if warehouse.Get(newPosition) == '#' {
			continue
		}

		if warehouse.Get(newPosition) == '.' {
			warehouse.Set(currentPosition, '.')
			warehouse.Set(newPosition, '@')
			currentPosition = newPosition
			continue
		}

		if warehouse.Get(newPosition) == '[' {
			ableToMove := canMove(newPosition, newPosition.Add(grid.RIGHT), dP, warehouse, true)
			if ableToMove {
				canMove(newPosition, newPosition.Add(grid.RIGHT), dP, warehouse, false)
				warehouse.Set(currentPosition, '.')
				warehouse.Set(newPosition, '@')
				currentPosition = newPosition
			}
		} else if warehouse.Get(newPosition) == ']' {
			ableToMove := canMove(newPosition.Add(grid.LEFT), newPosition, dP, warehouse, true)
			if ableToMove {
				canMove(newPosition.Add(grid.LEFT), newPosition, dP, warehouse, false)
				warehouse.Set(currentPosition, '.')
				warehouse.Set(newPosition, '@')
				currentPosition = newPosition
			}
		}
	}

	result := 0
	for idx := range warehouse {
		for jdx := range warehouse[idx] {
			if warehouse[idx][jdx] == '[' {
				result += 100*idx + jdx
			}
		}
	}
	fmt.Println(result)
}

func canMove(
	p1 grid.Point,
	p2 grid.Point,
	direction grid.Point,
	warehouse grid.Grid[byte],
	dryRun bool,
) bool {
	if warehouse.Get(p1) == '#' || warehouse.Get(p2) == '#' {
		return false
	}

	newP1 := p1.Add(direction)
	newP2 := p2.Add(direction)
	p1Intersect := warehouse.Get(newP1)
	p2Intersect := warehouse.Get(newP2)

	if direction == grid.UP || direction == grid.DOWN {
		if p1Intersect == '[' && p2Intersect == ']' && canMove(newP1, newP2, direction, warehouse, dryRun) {
			if !dryRun {
				warehouse.Set(p1, '.')
				warehouse.Set(newP1, '[')
				warehouse.Set(p2, '.')
				warehouse.Set(newP2, ']')
			}
			return true
		}

		p1Satisfied := p1Intersect == '.'
		if p1Intersect == ']' {
			p1Satisfied = canMove(newP1.Add(grid.LEFT), newP1, direction, warehouse, dryRun)
		}

		p2Satisfied := p2Intersect == '.'
		if p2Intersect == '[' {
			p2Satisfied = canMove(newP2, newP2.Add(grid.RIGHT), direction, warehouse, dryRun)
		}

		if p1Satisfied && p2Satisfied {
			if !dryRun {
				warehouse.Set(p1, '.')
				warehouse.Set(newP1, '[')
				warehouse.Set(p2, '.')
				warehouse.Set(newP2, ']')
			}
			return true
		}
	}

	if direction == grid.LEFT {
		if p1Intersect == '.' {
			if !dryRun {
				warehouse.Set(newP1, '[')
				warehouse.Set(p1, ']')
				warehouse.Set(p2, '.')
			}
			return true
		}

		if p1Intersect == '#' {
			return false
		}

		if p1Intersect == ']' && canMove(newP1.Add(grid.LEFT), newP1, direction, warehouse, dryRun) {
			if !dryRun {
				warehouse.Set(newP1, '[')
				warehouse.Set(p1, ']')
				warehouse.Set(p2, '.')
			}
			return true
		}
	}

	if direction == grid.RIGHT {
		if p2Intersect == '.' {
			if !dryRun {
				warehouse.Set(p1, '.')
				warehouse.Set(newP1, '[')
				warehouse.Set(newP2, ']')
			}
			return true
		}

		if p2Intersect == '#' {
			return false
		}

		if p2Intersect == '[' && canMove(newP2, newP2.Add(grid.RIGHT), direction, warehouse, dryRun) {
			if !dryRun {
				warehouse.Set(p1, '.')
				warehouse.Set(newP1, '[')
				warehouse.Set(newP2, ']')
			}
			return true
		}
	}

	return false
}

func dump(warehouse grid.Grid[byte]) {
	for idx := range warehouse {
		for jdx := range warehouse[idx] {
			if warehouse[idx][jdx] == '.' {
				fmt.Print(" ")
			} else if warehouse[idx][jdx] == '#' {
				fmt.Print("█")
			} else {
				fmt.Print(string(warehouse[idx][jdx]))
			}
		}
		fmt.Println()
	}
	fmt.Print(gocursor.MoveUp(len(warehouse)))
	fmt.Print(gocursor.MoveLeft(len(warehouse[0])))
}
