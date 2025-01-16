package TwentyOne

import (
	"math"
	"strings"

	"github.com/cdlewis/advent-of-code/2022/util"
)

type Node struct {
	Value     int
	Left      string
	Right     string
	Operation string
	IsValue   bool
}

func TwentyOne() int {
	raw := strings.Split(util.GetInput(21, false, `root: pppw + sjmn
	dbpl: 5
	cczh: sllz + lgvd
	zczc: 2
	ptdq: humn - dvpt
	dvpt: 3
	lfqf: 4
	humn: 5
	ljgn: 2
	sjmn: drzm * dbpl
	sllz: 4
	pppw: cczh / lfqf
	lgvd: ljgn * ptdq
	drzm: hmdt - zczc
	hmdt: 32`), "\n")

	graph := map[string]*Node{}

	for _, i := range raw {
		tokens := strings.Split(i, " ")

		name := tokens[0][:len(tokens[0])-1]

		if len(tokens) == 2 {
			graph[name] = &Node{Value: util.ToInt(tokens[1]), IsValue: true}
			continue
		}

		graph[name] = &Node{
			Left:      tokens[1],
			Operation: tokens[2],
			Right:     tokens[3],
		}
	}

	lower := 0
	upper := math.MaxInt / 100

	for lower < upper {
		mid := (lower + upper) / 2
		graph["humn"].Value = mid

		_, diff, _ := solve("root", graph)

		if diff > 0 {
			lower = mid + 1
		} else if diff < 0 {
			upper = mid - 1
		} else {
			return mid
		}
	}

	panic("No solution")
}

func solve(nodeName string, graph map[string]*Node) (int, int, bool) {
	node := graph[nodeName]

	if node.IsValue {
		return node.Value, 0, nodeName == "humn"
	}

	left, _, leftHuman := solve(node.Left, graph)
	right, _, rightHuman := solve(node.Right, graph)

	if nodeName == "root" {
		return left + right, left - right, true
	}

	result := 0
	switch node.Operation {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		result = left / right
	default:
		panic("unexpected")
	}

	humans := leftHuman || rightHuman || nodeName == "humn"

	// Re-write the graph to skip these nodes in future
	if !humans {
		graph[nodeName] = &Node{
			IsValue: true,
			Value:   result,
		}
	}

	return result, 0, humans
}
