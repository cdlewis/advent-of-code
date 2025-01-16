package three

import (
	"regexp"

	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/cast"
)

var testData = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
var validInstruction = regexp.MustCompile(`(mul\(([0-9]{1,3}),([0-9]{1,3})\))|(do\(\))|(don't\(\))`)

func Three() int {
	instructions := aoc.GetInput(3, false, testData)
	valid := validInstruction.FindAllStringSubmatch(instructions, -1)

	sum := 0
	ignore := false
	for _, v := range valid {
		prefix := v[0][:3]

		if prefix == "mul" {
			if !ignore {
				sum += (cast.ToInt(v[2]) * cast.ToInt(v[3]))
			}

			continue
		}

		ignore = prefix == "don"
	}

	return sum
}
