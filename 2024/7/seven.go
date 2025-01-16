package seven

import (
	"strconv"
	"strings"

	"github.com/cdlewis/advent-of-code/2024/util"
	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/cast"
)

type Equation struct {
	Target int
	Values []int
}

func Seven() int {
	rawEquations := strings.Split(aoc.GetInput(7, false, ""), "\n")

	result := 0
	for _, e := range rawEquations {
		tokens := strings.Split(e, " ")
		equation := Equation{
			Target: cast.ToInt(strings.TrimSuffix(tokens[0], ":")),
			Values: util.Map(tokens[1:], cast.ToInt[string]),
		}

		if isValid(equation) {
			result += equation.Target
		}
	}

	return result
}

var operations = [](func(x, y int) int){
	func(x, y int) int { return x + y },
	func(x, y int) int { return x * y },
	func(x, y int) int {
		return cast.ToInt(strconv.Itoa(x) + strconv.Itoa(y))
	},
}

func isValid(e Equation) bool {
	if len(e.Values) == 0 {
		return e.Target == 0
	}

	if len(e.Values) == 1 {
		return e.Values[0] == e.Target
	}

	for _, o := range operations {
		result := o(e.Values[0], e.Values[1])

		tryOperation := result <= e.Target && isValid(Equation{
			Values: append([]int{result}, e.Values[2:]...),
			Target: e.Target,
		})

		if tryOperation {
			return true
		}
	}

	return false
}
