package six

import (
	"fmt"

	"github.com/cdlewis/advent-of-code/2022/util"
)

func Six(useFallback bool, fallback string) int {
	raw := util.GetInput(6, useFallback, fallback)
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
			return i + 1
		}
	}

	panic("Nothing found")
}
