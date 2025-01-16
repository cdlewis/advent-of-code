package util

type List struct {
	Head *Node
	Tail *Node
}

func (l *List) Append(i int) {
	if l.Head == nil {
		l.Head = &Node{Val: i}
		l.Tail = l.Head
	} else {
		l.Tail = l.Tail.InsertAfter(i)
	}
}

func (l *List) Find(i int) *Node {
	return l.Head.Find(i)
}

func NewList() List {
	return List{}
}

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func (n *Node) Find(needle int) *Node {
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

func (n *Node) Idx(i int) *Node {
	curr := n

	for j := 0; j < i; j++ {
		curr = curr.Next
	}

	return curr
}

func (current *Node) MoveRight() {
	iPrev := current.Prev
	iNext := current.Next
	iNextNext := current.Next.Next

	current.Next = iNextNext
	current.Prev = iNext

	iNextNext.Prev = current

	iNext.Next = current
	iNext.Prev = iPrev

	iPrev.Next = iNext
}

func (current *Node) MoveLeft() {
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

func (n *Node) ToSlice() []int {
	curr := n.Next
	results := []int{n.Val}

	for curr != n {
		results = append(results, curr.Val)
		curr = curr.Next
	}

	return results
}

func (n *Node) InsertAfter(i int) *Node {
	next := &Node{Val: i}
	next.Prev = n
	n.Next = next

	n = next

	return next
}
