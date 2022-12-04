package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("./input")
	raw := strings.Split(string(dat), "\n")

	current := 0
	results := []int{}

	for _, c := range raw {
		if c == "" {
			results = append(results, current)
			current = 0
		}

		i, _ := strconv.ParseInt(c, 10, 64)

		current += int(i)
	}

	sort.Ints(results)

	last := len(results) - 1

	result := results[last] + results[last-1] + results[last-2]

	fmt.Println(result)

	if result != 205615 {
		panic("unexpected result")
	}
}
