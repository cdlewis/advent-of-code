package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/cast"
)

func main() {
	rawInput := strings.Split(aoc.GetInput(17, false, ""), "\n\n")
	program := cast.FindAllInt(rawInput[1])

	candidates := []int{0}
	for idx := len(program) - 1; idx >= 0; idx-- {
		var newCandidates []int
		for _, c := range candidates {
			for j := range 8 {
				testInput := (c << 3) + j
				if run(program, [3]int{testInput, 0, 0}, program[idx:]) {
					newCandidates = append(newCandidates, testInput)
				}
			}
		}
		candidates = newCandidates

		if len(candidates) == 0 {
			panic("NO MATCH")
		}
	}

	fmt.Println(slices.Min(candidates))
}

func run(program []int, registers [3]int, want []int) bool {
	instPtr := 0

	var output []int
	for instPtr < len(program) {
		opcode, operand := program[instPtr], program[instPtr+1]

		switch opcode {
		case 0: // adv
			registers[0] = registers[0] / (1 << getCombo(operand, registers))
		case 1: // bxl
			registers[1] = registers[1] ^ operand
		case 2: // bst
			registers[1] = getCombo(operand, registers) % 8
		case 3: // jnz
			if registers[0] != 0 {
				instPtr = operand
				continue
			}
		case 4: // bxc
			registers[1] = registers[1] ^ registers[2]
		case 5: // out
			output = append(output, getCombo(operand, registers)%8)
			if len(output) > len(want) || output[len(output)-1] != want[len(output)-1] {
				return false
			}
		case 6: // bdv
			registers[1] = registers[0] / (1 << getCombo(operand, registers))
		case 7: // cdv
			registers[2] = registers[0] / (1 << getCombo(operand, registers))
		}

		instPtr += 2
	}

	if len(want) != len(output) {
		return false
	}

	for i := range want {
		if want[i] != output[i] {
			return false
		}
	}

	return true
}

func getCombo(operand int, registers [3]int) int {
	if operand <= 3 {
		return operand
	}

	if operand < 7 {
		return registers[operand-4]
	}

	panic("invalid operand")
}
