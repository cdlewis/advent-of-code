package TwentyOne

import (
	"math"
	"strings"

	"github.com/cdlewis/advent-of-code/util"
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

	// Lazy 'binary' search for a starting point somewhat close to the result

	i := 0
	skipAmount := math.MaxInt / 100
	for skipAmount > 1 {
		graph["humn"].Value = i + skipAmount

		if _, diff, _ := solve("root", graph); diff > 0 {
			i = i + skipAmount
			continue
		}

		graph["humn"].Value = i - skipAmount
		skipAmount /= 2
	}

	// Brute force the remainder

	for ; i < math.MaxInt; i++ {
		graph["humn"].Value = i

		_, diff, _ := solve("root", graph)

		if diff == 0 {
			return i
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
