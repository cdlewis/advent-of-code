package main

import (
	"fmt"
	"strings"

	"github.com/cdlewis/advent-of-code/util"
)

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func (n *Node) find(needle int) *Node {
	start := n
	curr := n

	if n.Val == needle {
		return n
	}

	curr = start.Next

	for curr != start {
		if curr.Val == needle {
			return curr
		}

		curr = curr.Next
	}

	panic("not found")
}

func (n *Node) toSlice() []int {
	curr := n.Next
	results := []int{n.Val}

	for curr != n {
		results = append(results, curr)
		curr = curr.Next
	}

	return results
}

func main() {
	raw := util.Map(strings.Split(util.GetInput(20, true, `1
	2
	-3
	3
	-2
	0
	4`), "\n"), util.ToInt[string])

	var root *Node
	var curr *Node

	for idx, i := range raw {
		if idx == 0 {
			root := &Node{Val: i}
			curr = root
			continue
		}

		next := &Node{Val: i}
		next.Prev = curr
		curr.Next = next

		curr = next
	}

	// join the start/end
	root.Prev = curr
	curr.Next = root

	for _, i := range raw {
		if i == 0 {
			continue
		}

		current := root.find(i)

		for j := 0; j < i; j++ {
			// forward
			if i > 0 {
				iPrev := current.Prev
				iNext := current.Next
				iNextNext := current.Next.Next

				current.Next = iNextNext
				current.Prev = iNext

				iNextNext.Prev = current

				iNext.Next = current
				iNext.Prev = iPrev

				iPrev.Next = iNext
			} else { // backward
				iPrevPrev := current.Prev.Prev
				iPrev := current.Prev
				iNext := current.Next

				iPrevPrev.Next = current

				current.Prev = iPrevPrev
				current.Next = iPrev

				iPrev.Prev = current
				iPrev.Next = iNext

				iNext.Prev = iPrev
			}
		}

		fmt.Println(root.toSlice())
	}

	c := root.toSlice()

	zeroIdx := 0
	for idx, i := range c {
		if i == 0 {
			zeroIdx = idx
		}
	}

	one := c[(zeroIdx+1000)%(len(raw))]
	two := c[(zeroIdx+2000)%(len(raw))]
	three := c[(zeroIdx+3000)%(len(raw))]

	fmt.Println("one", one, (zeroIdx+1000)%(len(raw)), "two", two, (zeroIdx+2000)%(len(raw)), "three", three)
	fmt.Println(one + two + three)
}

func without(s []int, i int) []int {
	result := []int{}

	for idx, v := range s {
		if idx == i {
			continue
		}

		result = append(result, v)
	}

	return result
}

func insertAfter(s []int, toInsert, target int) []int {
	result := []int{}

	for _, j := range s {
		if j == target {
			result = append(result, target, toInsert)
		} else {
			result = append(result, j)
		}
	}

	return result
}

func insertBefore(s []int, toInsert, target int) []int {
	result := []int{}

	for idx, j := range s {
		if j == target {

			if idx == 0 {
				return append(s, toInsert)
			}

			result = append(result, toInsert, target)
		} else {
			result = append(result, j)
		}
	}

	return result
}

func find(s []int, i int) int {
	for jdx, j := range s {
		if j == i {
			return jdx
		}
	}
	panic("not found")
}
