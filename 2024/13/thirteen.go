package thirteen

import (
	"regexp"

	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/cast"
	"github.com/cdlewis/advent-of-code/2024/util/grid"
)

const padding = 10_000_000_000_000

var re = regexp.MustCompile("[0-9]+")

type Game struct {
	A     grid.Point
	B     grid.Point
	Prize grid.Point
}

func Thirteen() int {
	input := aoc.GetInput(13, false, "")
	rawNumbers := re.FindAllString(input, -1)

	total := 0
	for i := 0; i+5 < len(rawNumbers); i = i + 6 {
		game := Game{
			A:     grid.Point{cast.ToInt(rawNumbers[i]), cast.ToInt(rawNumbers[i+1])},
			B:     grid.Point{cast.ToInt(rawNumbers[i+2]), cast.ToInt(rawNumbers[i+3])},
			Prize: grid.Point{cast.ToInt(rawNumbers[i+4]) + padding, cast.ToInt(rawNumbers[i+5]) + padding},
		}

		cost, found := cheapestRoute(game)
		if found {
			total += cost
		}
	}

	return total
}

func cheapestRoute(rules Game) (int, bool) {
	// cramer's rule
	numA := determinent(rules.Prize, rules.B) / determinent(rules.A, rules.B)
	numB := determinent(rules.A, rules.Prize) / determinent(rules.A, rules.B)

	if rules.A[0]*numA+rules.B[0]*numB == rules.Prize[0] && rules.A[1]*numA+rules.B[1]*numB == rules.Prize[1] {
		return 3*numA + numB, true
	}

	return -1, false
}

// treat col1 and col2 as forming a 2x2 matrix and calculate determinant
func determinent(col1, col2 grid.Point) int {
	return col1[0]*col2[1] - col1[1]*col2[0]
}
