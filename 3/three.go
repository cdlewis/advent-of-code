package main

import (
	"fmt"
	"os"
	"strings"
)

func calculateScore(item rune) int {
	i := byte(item)
	if i >= 'a' && i <= 'z' {
		return int(i-'a') + 1
	}

	return int(i - 'A' + 27)
}

func keyIntersection(maps []map[rune]bool) rune {
	keyCounts := map[rune]int{}
	for _, inputMap := range maps {
		for key := range inputMap {
			keyCounts[key]++
		}
	}

	for key, count := range keyCounts {
		if count == len(maps) {
			return key
		}
	}

	panic("no key found")
}

func main() {
	dat, _ := os.ReadFile("./input")

	raw := strings.Split(string(dat), "\n")
	total := 0

	for i := 0; i < len(raw); i += 3 {
		seenItemsPerBackpack := make([]map[rune]bool, 3)

		for j := i; j < i+3; j++ {
			backpack := raw[j]

			for _, item := range backpack {
				if seenItemsPerBackpack[j-i] == nil {
					seenItemsPerBackpack[j-i] = map[rune]bool{}
				}
				seenItemsPerBackpack[j-i][item] = true
			}
		}

		item := keyIntersection(seenItemsPerBackpack)
		result := calculateScore(item)

		total += result
	}

	if total != 2650 {
		panic("Incorrect answer")
	}

	fmt.Println("Total", total)
}
