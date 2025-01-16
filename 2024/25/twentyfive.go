package main

import (
	"fmt"
	"strings"

	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/grid"
)

func main() {
	rawPlans := strings.Split(aoc.GetInput(25, false, ""), "\n\n")

	var locks [][5]int
	var keys [][5]int

	for _, i := range rawPlans {
		lines := strings.Split(i, "\n")
		schematic := grid.ToByteGrid(i)

		result := [5]int{}
		if lines[0] == "#####" {
			for j := 0; j < 5; j++ {
				size := 0
				for i := 1; i < len(lines); i++ {
					if schematic[i][j] == '.' {
						break
					}
					size++
				}
				result[j] = size
			}
			locks = append(locks, result)
		} else {
			for j := 0; j < 5; j++ {
				size := 0
				for i := len(lines) - 2; i >= 0; i-- {
					if schematic[i][j] == '.' {
						break
					}
					size++
				}
				result[j] = size
			}
			keys = append(keys, result)
		}
	}

	safe := 0
	for _, k := range keys {
	SEARCH_LOCKS:
		for _, l := range locks {
			for idx := range k {
				if l[idx]+k[idx] > 5 {
					continue SEARCH_LOCKS
				}
			}
			safe++
		}
	}

	fmt.Println(safe)
}
