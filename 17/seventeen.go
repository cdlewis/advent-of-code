package seventeen

import (
	"github.com/cdlewis/advent-of-code/util"
)

func Seventeen(input string) int {
	caveHeight := 100_000
	c := make([][]int, caveHeight)
	for i := range c {
		c[i] = make([]int, 7)
	}

	gameState := GameState{
		Cave:       c,
		Towersize:  0,
		ShapeIndex: 1,
		WindIndex:  0,
	}

	cycleChecker := NewCycleChecker()
	boostedHeight := 0
	currentShape := append([][2]int{}, shapes[0]...)
	for idx := range currentShape {
		currentShape[idx][0] += 3
		currentShape[idx][1] += 2
	}

	target := 1_000_000_000_000
	for gameState.ShapesStopped <= target {
		windDirection := -1
		if input[gameState.WindIndex] == '>' {
			windDirection = 1
		}

		shouldMoveWind := true
		for _, p := range currentShape {
			if !util.ValidCoordinate(p[0], p[1]+windDirection, c) || c[p[0]][p[1]+windDirection] == 2 {
				shouldMoveWind = false
				break
			}
		}

		if shouldMoveWind {
			for idx := range currentShape {
				currentShape[idx][1] += windDirection
			}
		}

		shouldMoveDown := true
		for _, p := range currentShape {
			if !util.ValidCoordinate(p[0]-1, p[1], c) || c[p[0]-1][p[1]] == 2 {
				shouldMoveDown = false
				break
			}
		}

		if shouldMoveDown {
			for idx := range currentShape {
				currentShape[idx][0] -= 1
			}
		}

		gameState.WindIndex = (gameState.WindIndex + 1) % len(input)

		if !shouldMoveDown {
			gameState.ShapesStopped++

			// write shape into stone
			for _, p := range currentShape {
				gameState.Cave[p[0]][p[1]] = 2

				// update max height as we go
				gameState.Towersize = util.Max(gameState.Towersize, p[0])
			}

			if gameState.ShapesStopped == target {
				return gameState.Towersize + boostedHeight + 1
			}

			if isCycle, shapeInterval, heightInterval := cycleChecker.IsCycle(gameState); isCycle {
				boosts := (target - gameState.ShapesStopped) / shapeInterval
				gameState.ShapesStopped = gameState.ShapesStopped + boosts*shapeInterval
				boostedHeight = boostedHeight + boosts*heightInterval
			}

			// generate new rock
			currentShape = append([][2]int{}, shapes[gameState.ShapeIndex]...)
			gameState.ShapeIndex = (gameState.ShapeIndex + 1) % len(shapes)

			// translate shape to starting pos
			for idx := range currentShape {
				currentShape[idx][0] += (gameState.Towersize + 4)
				currentShape[idx][1] += 2
			}
		}
	}

	return gameState.Towersize + boostedHeight + 1
}

const gameStateSize = 98

type CycleChecker struct {
	checkPoints map[[gameStateSize + 2]int]GameState
}

func NewCycleChecker() CycleChecker {
	return CycleChecker{checkPoints: map[[gameStateSize + 2]int]GameState{}}
}

func (c *CycleChecker) IsCycle(g GameState) (bool, int, int) {
	rows := gameStateSize / 7
	if g.Towersize < rows+1 {
		return false, 0, 0
	}

	heights := [gameStateSize + 2]int{}
	for i := g.Towersize; i > g.Towersize-rows; i-- {
		for j := 0; j < 7; j++ {
			heights[(g.Towersize-i)*7+j] = g.Cave[i][j]
		}
	}

	for i := g.Towersize; i > g.Towersize-rows; i-- {
		for j := 0; j < 7; j++ {
			heights[(g.Towersize-i)*7+j] = g.Cave[i][j]
		}
	}

	heights[gameStateSize] = g.WindIndex
	heights[gameStateSize+1] = g.ShapeIndex % len(shapes)

	if prevState, ok := c.checkPoints[heights]; ok {
		shapesStoppedFromCycle := g.ShapesStopped - prevState.ShapesStopped
		towerHeightFromCycle := g.Towersize - prevState.Towersize

		return true, shapesStoppedFromCycle, towerHeightFromCycle
	}

	c.checkPoints[heights] = g

	return false, 0, 0

}

type GameState struct {
	Cave          [][]int
	Towersize     int
	ShapeIndex    int
	WindIndex     int
	ShapesStopped int
}

var shapes = [][][2]int{
	{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
	},
	{
		{0, 1},
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 1},
	},
	{
		{2, 2},
		{1, 2},
		{0, 0},
		{0, 1},
		{0, 2},
	},
	{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
	},
	{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	},
}
