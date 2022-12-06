package two

import (
	"strings"

	"github.com/cdlewis/advent-of-code/util"
)

const LOSE = 0
const DRAW = 1
const WIN = 2

const ROCK = 0
const PAPER = 1
const SCISSORS = 2

var reward = []int{0, 3, 6}

var outcomesMatrix = [][]int{
	{DRAW, LOSE, WIN}, // ROCK
	{WIN, DRAW, LOSE}, // PAPER
	{LOSE, WIN, DRAW}, // SCISSORS
}

func Two() int {
	raw := strings.Split(util.GetInput(2, false, ""), "\n")

	score := 0

	for _, game := range raw {
		oponentMove := int(game[0] - 'A')
		desiredOutcome := int(game[2] - 'X')

		myMove := -1
		for move, outcome := range outcomesMatrix {
			if outcome[oponentMove] == desiredOutcome {
				myMove = move
				break
			}
		}
		if myMove == -1 {
			panic("invariant violation: no valid move picked")
		}

		score += (reward[desiredOutcome] + myMove + 1)
	}

	return score
}
