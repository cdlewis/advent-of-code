package twenty

import (
	"strings"

	"github.com/cdlewis/advent-of-code/2022/util"
)

var decryptionKey = 811589153

func Twenty() int {
	raw := util.Map(strings.Split(util.GetInput(20, false, `1
	2
	-3
	3
	-2
	0
	4`), "\n"), util.ToInt[string])

	list := util.NewList()
	list.Append(raw[0] * decryptionKey)
	orderedNodes := []*util.Node{list.Head}

	for _, i := range raw[1:] {
		list.Append(i * decryptionKey)
		orderedNodes = append(orderedNodes, list.Tail)
	}

	// join the start/end
	list.Head.Prev, list.Tail.Next = list.Tail, list.Head

	for n := 0; n < 10; n++ {
		for _, current := range orderedNodes {
			if current.Val == 0 {
				continue
			}

			for j := 0; j < util.Abs(current.Val)%(len(raw)-1); j++ {
				if current.Val > 0 {
					current.MoveRight()
				} else {
					current.MoveLeft()
				}
			}

		}
	}

	zeroNode := list.Find(0)

	one := zeroNode.Idx(1000)
	two := one.Idx(1000)
	three := two.Idx(1000)

	return one.Val + two.Val + three.Val
}
