package main

import (
	"fmt"

	"github.com/cdlewis/advent-of-code/util"
)

func main() {
	raw := util.GetInput(6, true, "mjqjpqmgbljsphdztnvjfqwrcgsmlb")
	windowSize := 14

	seen := map[byte]int{}
	for i := 0; i < len(raw); i++ {
		seen[raw[i]]++

		if i >= windowSize {
			seen[raw[i-windowSize]]--
		}

		isUnique := true
		for _, v := range seen {
			if v > 1 {
				isUnique = false
				break
			}
		}

		if i >= windowSize && isUnique {
			fmt.Println("Found", i+1)
			return
		}
	}

	panic("Nothing found")
}
