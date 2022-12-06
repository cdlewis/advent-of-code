package main

import (
	"fmt"
	"os"
)

func main() {
	dat, _ := os.ReadFile("./input")
	raw := string(dat)
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
