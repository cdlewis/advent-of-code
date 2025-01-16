package main

import (
	"fmt"
	"maps"
	"math"
	"strings"

	"github.com/cdlewis/advent-of-code/2024/util"
	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/cast"
)

const _magic = 16777216

func main() {
	input := util.Map(strings.Split(aoc.GetInput(22, false, ""), "\n"), cast.ToInt)

	var changeMaps []map[[4]int]int
	allChanges := map[[4]int]int{}
	for _, i := range input {
		changes := getRand(i, 2000)
		changeMaps = append(changeMaps, changes)
		maps.Copy(allChanges, changes)
	}

	best := 0

	for change := range allChanges {
		total := 0
		for _, changeMap := range changeMaps {
			total += changeMap[change]
		}

		best = max(total, best)
	}

	fmt.Println(best)
}

func getRand(seed int, iterations int) map[[4]int]int {
	nums := make([]int, 0, 2000)
	nums = append(nums, seed%10)

	for range iterations {
		seed = (seed ^ (seed * 64)) % _magic
		seed = seed ^ int(math.Floor(float64(seed)/float64(32)))%_magic
		seed = (seed ^ (seed * 2048)) % _magic
		nums = append(nums, seed%10)
	}

	seqToPrice := map[[4]int]int{}
	for i := 3; i < len(nums); i++ {
		base := 0
		if i > 3 {
			base = nums[i-4]
		}

		changeKey := [4]int{
			nums[i-3] - base,
			nums[i-2] - nums[i-3],
			nums[i-1] - nums[i-2],
			nums[i] - nums[i-1],
		}

		if _, ok := seqToPrice[changeKey]; ok {
			continue
		}
		seqToPrice[changeKey] = nums[i]
	}

	return seqToPrice
}
