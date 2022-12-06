package three

import (
	"strings"

	"github.com/cdlewis/advent-of-code/util"
)

func calculateScore(item rune) int {
	i := byte(item)
	if i >= 'a' && i <= 'z' {
		return int(i-'a') + 1
	}

	return int(i - 'A' + 27)
}

func Three() int {
	raw := strings.Split(util.GetInput(3, false, ""), "\n")
	total := 0

	for i := 0; i < len(raw); i += 3 {
		item := util.IntersectionString(raw[i : i+3]...)
		result := calculateScore(rune(item[0]))

		total += result
	}

	return total
}
