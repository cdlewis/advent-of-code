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
	x := 1  // cursor position
	r1 := 0 // register 1

	// Clear the screen and create some room
	fmt.Print("\033[2J")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n")

	for len(serialisedInstructions) > 0 {
		// Move cursor to the top of the screen and write our summary for this cycle
		fmt.Printf("\033[%d;%dH", 0, 0)
		fmt.Println("Cycle:       ", cycle)
		fmt.Println("Instruction: ", serialisedInstructions[0], "               ")

		currentRow := 0
		currentColumn := cycle
		currentRow = cycle / 40
		currentColumn = cycle % (util.Max(currentRow, 1) * 40)

		// Move cursor to the current row/column (with some spacing to separate from the summary)
		fmt.Printf("\033[%d;%dH", currentRow+5, currentColumn)

		if currentColumn >= x-1 && currentColumn <= x+1 {
			fmt.Print("#")
		}

		t := strings.Split(serialisedInstructions[0], " ")
		serialisedInstructions = serialisedInstructions[1:]

		if t[0] == "addx" {
			// Invent a new instruction (add register r1 to counter) to simulate delaying execution addx
			serialisedInstructions = append([]string{"add_r1"}, serialisedInstructions...)
			r1 = util.ToInt(t[1])
		} else if t[0] == "add_r1" {
			x += r1
		}

		cycle++

		time.Sleep(10 * time.Millisecond)
	}
}
