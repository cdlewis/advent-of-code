package grid

import (
	"container/heap"
	"math"
)

type WeightedPoint struct {
	To     Point
	Weight int
}

type EdgeFetcher func(p Point) []WeightedPoint

type Item struct {
	Node     Point
	Distance int
	Index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.Index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Update(item *Item, distance int) {
	item.Distance = distance
	heap.Fix(pq, item.Index)
}

func Dijkstra(start, target Point, fetchEdges EdgeFetcher) (int, map[Point]Point) {
	distances := make(map[Point]int)
	prev := make(map[Point]Point)
	pq := make(PriorityQueue, 0)

	// Initialize distances and priority queue
	for _, node := range []Point{start} {
		distances[node] = math.MaxInt
	}
	distances[start] = 0

	heap.Push(&pq, &Item{Node: start, Distance: 0})

	// Main loop
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item).Node

		// Stop early if we reached the target
		if current == target {
			break
		}

		// Fetch neighbors and update distances
		for _, edge := range fetchEdges(current) {
			alt := distances[current] + edge.Weight
			if alt < distances[edge.To] {
				distances[edge.To] = alt
				prev[edge.To] = current
				heap.Push(&pq, &Item{Node: edge.To, Distance: alt})
			}
		}
	}

	return distances[target], prev
}
