package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/cdlewis/advent-of-code/util"
)

func main() {
	serialisedInstructions := strings.Split(util.GetInput(10, false, ""), "\n")

	cycle := 0
	x := 1 // cursor position

	// Clear the screen and create some room
	fmt.Print("\033[2J")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n")

	for len(serialisedInstructions) > 0 {
		// Move cursor to the top of the screen and write our summary for this cycle
		fmt.Printf("\033[%d;%dH", 0, 0)
		fmt.Print("\u001b[1000D")
		fmt.Println("Cycle:       ", cycle)
		fmt.Print("\u001b[1000D")
		fmt.Println("Instruction: ", serialisedInstructions[0])

		currentRow := 0
		currentColumn := cycle
		if cycle/40 > 0 { // avoid div by 0 issues
			currentRow = cycle / 40
			currentColumn = cycle % (currentRow * 40)
		}

		// Move cursor to the current row/column (with some spacing to separate from the summary)
		fmt.Printf("\033[%d;%dH", currentRow+5, currentColumn)

		if currentColumn >= x-1 && currentColumn <= x+1 {
			fmt.Print("#")
		}

		t := strings.Split(serialisedInstructions[0], " ")
		serialisedInstructions = serialisedInstructions[1:]

		if t[0] == "addx" {
			// Invent a new instruction to simulate delaying execution addx
			serialisedInstructions = append([]string{"do_addx " + t[1]}, serialisedInstructions...)
		} else if t[0] == "do_addx" {
			x += util.ToInt(t[1])
		}

		cycle++
		time.Sleep(100 * time.Millisecond)
	}
}
