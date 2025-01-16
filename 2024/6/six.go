package six

import (
	"runtime"
	"sync"

	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/grid"
)

var playerStates = map[byte][2]int{
	'^': {-1, 0},
	'v': {1, 0},
	'<': {0, -1},
	'>': {0, 1},
}

var nextState = map[[2]int][2]int{
	{-1, 0}: {0, 1},
	{0, 1}:  {1, 0},
	{1, 0}:  {0, -1},
	{0, -1}: {-1, 0},
}

type GuardState struct {
	Position  grid.Point
	Direction grid.Point
}

func (g GuardState) Serialize() [4]int {
	return [4]int{g.Direction[0], g.Direction[1], g.Position[0], g.Position[1]}
}

func Six() int {
	grid := grid.ToByteGrid(aoc.GetInput(6, false, ""))

	startingPosition := getInitialState(grid)
	chunks := runtime.NumCPU()
	chunkSize := len(grid) / chunks
	results := make([]int, chunks)
	var wg sync.WaitGroup
	for i := range chunks {
		wg.Add(1)
		go func() {
			results[i] = search(
				startingPosition,
				i*chunkSize,
				i*chunkSize+chunkSize,
				grid,
			)
			wg.Done()
		}()
	}

	wg.Wait()

	locations := 0
	for _, i := range results {
		locations += i
	}

	return locations
}

func search(
	startingPosition GuardState,
	fromI int,
	toI int,
	grid [][]byte,
) int {
	locations := 0

	for idx := fromI; idx < toI; idx++ {
		for jdx, j := range grid[idx] {
			if idx == startingPosition.Position[0] && jdx == startingPosition.Position[1] {
				continue
			}

			if j == '#' {
				continue
			}

			currentPosition := startingPosition
			if hasCycle(currentPosition, grid, idx, jdx) {
				locations++
			}
		}
	}

	return locations
}

func getInitialState(grid [][]byte) GuardState {
	for idx, i := range grid {
		for jdx, j := range i {
			if state, ok := playerStates[j]; ok {
				return GuardState{
					Direction: state,
					Position:  [2]int{idx, jdx},
				}
			}
		}
	}

	panic("no guard present")
}

func hasCycle(
	state GuardState,
	floorMap grid.Grid[byte],
	newBarrierI int,
	newBarrierJ int,
) bool {
	seen := map[[4]int]struct{}{}

	for {
		serializedState := state.Serialize()
		if _, ok := seen[serializedState]; ok {
			return true
		}

		seen[serializedState] = struct{}{}

		newPoint := state.Position.Add(state.Direction)
		if !floorMap.ValidPoint(newPoint) {
			return false
		}

		if (newPoint[0] == newBarrierI && newPoint[1] == newBarrierJ) || floorMap.Get(newPoint) == '#' {
			state.Direction = nextState[state.Direction]
		} else {
			state.Position = newPoint
		}
	}
}
