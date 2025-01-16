package main

import (
	"fmt"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/cast"
)

var testInput = `x00: 0
x01: 1
x02: 0
x03: 1
x04: 0
x05: 1
y00: 0
y01: 0
y02: 1
y03: 1
y04: 0
y05: 1

x00 AND y00 -> z05
x01 AND y01 -> z02
x02 AND y02 -> z01
x03 AND y03 -> z03
x04 AND y04 -> z04
x05 AND y05 -> z00`

type Node struct {
	Operation string
	Left      string
	Right     string
}

func main() {
	input := strings.Split(aoc.GetInput(24, false, testInput), "\n\n")

	mapGraph := map[string]Node{}

	wireGroups := map[byte][]string{}
	baseResolvedNodes := map[string]int{}

	for _, i := range strings.Split(input[0], "\n") {
		tokens := strings.Split(i, ": ")
		baseResolvedNodes[tokens[0]] = cast.ToInt(tokens[1])
		wireGroups[tokens[0][0]] = append(wireGroups[tokens[0][0]], tokens[0])
	}

	for _, i := range strings.Split(input[1], "\n") {
		tokens := strings.Split(i, " ")
		mapGraph[tokens[4]] = Node{
			Operation: tokens[1],
			Left:      tokens[0],
			Right:     tokens[2],
		}
		wireGroups[tokens[2][0]] = append(wireGroups[tokens[2][0]], tokens[2])
		wireGroups[tokens[4][0]] = append(wireGroups[tokens[4][0]], tokens[4])
	}

	for k := range wireGroups {
		slices.Sort(wireGroups[k])
		slices.Reverse(wireGroups[k])
	}

	for _, node := range wireGroups['z'] {

		if node == "z00" || node == "z45" {
			continue
		}

		if assertShape(node, mapGraph, Shape{LeftName: "x", RightName: "y"}) {
			fmt.Println("BAD NODE", mapGraph[node], node)
			continue
		}

		if !assertShape(node, mapGraph, Shape{Operation: "XOR"}) {
			fmt.Println("BAD NODE", mapGraph[node], node)
			continue
		}

		match := assertShape(node, mapGraph, Shape{
			Operation: "XOR",
			Left: &Shape{
				Operation: "XOR",
				LeftName:  "x",
				RightName: "y",
			},
			Right: &Shape{
				Operation: "OR",
				Left: &Shape{
					Operation: "AND",
					LeftName:  "x",
					RightName: "y",
				},
				Right: &Shape{
					Operation: "AND",
					Left: &Shape{
						Operation: "XOR",
						LeftName:  "x",
						RightName: "y",
					},
				},
			},
		})

		if !match {
			fmt.Println(node, mapGraph[node])
		}
	}

	//score(wireGroups, mapGraph, baseResolvedNodes)
}

type Shape struct {
	Operation string
	LeftName  string
	Left      *Shape
	RightName string
	Right     *Shape
}

func assertShape(node string, graph map[string]Node, want Shape) bool {
	curr := graph[node]
	if want.Operation != "" && curr.Operation != want.Operation {
		return false
	}

	if want.LeftName != "" && !strings.HasPrefix(curr.Left, want.LeftName) && !strings.HasPrefix(curr.Right, want.LeftName) {
		return false
	}

	leftValid := true
	if want.Left != nil {
		leftValid = assertShape(curr.Left, graph, *want.Left) || assertShape(curr.Right, graph, *want.Left)
	}

	rightValid := true
	if want.Right != nil {
		leftValid = assertShape(curr.Right, graph, *want.Right) || assertShape(curr.Left, graph, *want.Right)
	}

	if leftValid && rightValid {
		return true
	}
	//fmt.Println(curr, want)
	return false
}

func badGates(g map[string]Node) []string {
	var badGates []string
	for name, node := range g {
		//fmt.Println(name, node)
		leftType := node.Left[0]
		rightType := node.Right[0]
		if name[0] == 'z' && node.Operation == "XOR" && ((leftType == 'x' && rightType == 'y') || (leftType == 'y' && rightType == 'z')) {
			fmt.Println("BAD", name)
			badGates = append(badGates, name)
		}

		if node.Operation != "XOR" && name[0] == 'z' && name != "z45" {
			fmt.Println("BAD", name, node)
			badGates = append(badGates, name)
		}
	}

	return badGates
}

func score(
	wireGroups map[byte][]string,
	graph map[string]Node,
	baseResolved map[string]int,
) int {
	score := 0
	got, expected, valid := evaluateGraph(wireGroups, graph, baseResolved)
	if !valid {
		return math.MaxInt
	}
	for bitNum := range wireGroups['z'] {
		if ((1 << bitNum) & got) == ((1 << bitNum) & expected) {
			continue
		}
		score++
		fmt.Println("Bitnum", bitNum, wireGroups['z'][len(wireGroups['z'])-bitNum-1])
	}
	fmt.Println(score)
	return score
}

/*
func findSatisfying(

	numBits int,
	wireGroups map[byte][]string,
	graph map[string]Node,
	baseResolved map[string]int,
	swapped [][2]string,

	) [][][2]string {
		used := map[string]bool{}
		for _, c := range swapped {
			used[c[0]] = true
			used[c[1]] = true
		}

		var results [][][2]string

		baseSwappable := findSwappableNodes(numBits, wireGroups, graph)
		swappable := make([]string, 0, len(baseSwappable))
		for _, i := range baseSwappable {
			if used[i] {
				continue
			}
			swappable = append(swappable, i)
		}
		fmt.Println("\t num swappable", len(swappable))

		candidateGroups := generatePairs(swappable)
		fmt.Println(fmt.Println("\t candidate groups", len(candidateGroups)))
		for _, group := range candidateGroups {
			group = append(group, swapped...)
			nGraph := maps.Clone(graph)
			for _, pair := range group {
				nGraph[pair[0]], nGraph[pair[1]] = nGraph[pair[1]], nGraph[pair[0]]
			}

			if evaluateGraph(numBits, wireGroups, nGraph, baseResolved) {
				results = append(results, group)
			}
		}

		return results
	}
*/

func evaluateGraph(wireGroups map[byte][]string, graph map[string]Node, seen map[string]int) (int, int, bool) {
	resolved := maps.Clone(seen)

	xString := ""
	for _, k := range wireGroups['x'] {
		nextX, ok := resolveNode(k, graph, resolved)
		if !ok {
			//fmt.Println("Cycle detected")
			return 0, 0, false
		}
		xString += strconv.Itoa(nextX)
	}
	xVal, _ := strconv.ParseInt(xString, 2, 64)

	yString := ""
	for _, k := range wireGroups['y'] {
		nextY, ok := resolveNode(k, graph, resolved)
		if !ok {
			//fmt.Println("Cycle detected")
			return 0, 0, false
		}
		yString += strconv.Itoa(nextY)
	}
	yVal, _ := strconv.ParseInt(yString, 2, 64)

	zString := ""
	for _, k := range wireGroups['z'] {
		nextZ, ok := resolveNode(k, graph, resolved)
		if !ok {
			//fmt.Println("Cycle detected")
			return 0, 0, false
		}
		zString += strconv.Itoa(nextZ)
	}
	zVal, _ := strconv.ParseInt(zString, 2, 64)

	//fmt.Printf("%b + %b = %b, correct is %b\n", xVal, yVal, zVal, xVal+yVal)
	return int(xVal) + int(yVal), int(zVal), true
}

func resolveNode(start string, graph map[string]Node, resolved map[string]int) (int, bool) {
	//fmt.Println(start)
	if result, ok := resolved[start]; ok {
		return result, result != -1
	}
	resolved[start] = -1

	if _, ok := graph[start]; !ok {
		fmt.Println(start)
		panic("invalid node")
	}

	leftSide, leftOk := resolveNode(graph[start].Left, graph, resolved)
	rightSide, rightOk := resolveNode(graph[start].Right, graph, resolved)
	result := operations[graph[start].Operation](leftSide, rightSide)

	if !leftOk || !rightOk {
		return -1, false
	}

	resolved[start] = result

	return result, true
}

var operations = map[string](func(x, y int) int){
	"XOR": func(x, y int) int { return x ^ y },
	"AND": func(x, y int) int { return x & y },
	"OR":  func(x, y int) int { return x | y },
}

/*
 100001100101000011000110111001110010110100011
 111110101011110100000111001101111111011001101

 1100000010100111001001101100111110001111110000, correct is
 1100000010000110111001110000111110010001110000
*/

// hkj XOR bmh -> z08
