package main

import (
	"fmt"
	"os"
	"strings"
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

func main() {
	dat, _ := os.ReadFile("./input")

	raw := strings.Split(string(dat), "\n")

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

	fmt.Println(score)

	if score != 9541 {
		panic("incorrect score")
	}
}
