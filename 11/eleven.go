package eleven

import (
	"sort"
	"strings"

	"github.com/cdlewis/advent-of-code/util"
)

type Monkey struct {
	items       []int
	operation   []string
	test        int
	ifTrue      int
	ifFalse     int
	inspections int
}

func Eleven(useTest bool, testString string) int {
	raw := util.GetInput(11, useTest, testString)
	tokenisedResponses := util.Filter(util.Map(strings.Split(raw, "\n"), func(i string) []string {
		return util.Filter(strings.Split(strings.TrimSpace(i), " "), func(j string) bool {
			return len(j) > 0
		})
	}), func(k []string) bool { return len(k) > 0 })

	monkeys := []Monkey{}
	currentMonkey := Monkey{}

	// Build a slice of Monkeys from the input

	for idx, tokens := range tokenisedResponses {
		switch tokens[0] {
		case "Monkey":
			if idx > 0 {
				monkeys = append(monkeys, currentMonkey)
				currentMonkey = Monkey{}
			}
		case "Starting":
			currentMonkey.items = util.Map(tokens[2:], func(t string) int {
				return int(util.ToInt(strings.ReplaceAll(t, ",", "")))
			})
		case "Operation:":
			currentMonkey.operation = tokens[3:]
		case "Test:":
			currentMonkey.test = util.ToInt(tokens[len(tokens)-1])
		case "If":
			if tokens[1] == "true:" {
				currentMonkey.ifTrue = util.ToInt(tokens[len(tokens)-1])
			} else {
				currentMonkey.ifFalse = util.ToInt(tokens[len(tokens)-1])
			}
		default:
			panic("Unexpected token: " + tokens[0])
		}
	}
	monkeys = append(monkeys, currentMonkey) // don't forget the last monkey!

	// Create an upper bound for worry values to stop them overflowing. The upper bound
	// will be the product of the test values. Modulus is transitive over addition and
	// multiplication so this safely retains the information we need to calculate remainders.
	upperBound := util.Reduce(monkeys, func(acc int, m Monkey) int { return m.test * acc }, 1)

	for round := 0; round < 10000; round++ {
		for idx := range monkeys {
			for len(monkeys[idx].items) > 0 {
				currentItemWorry := monkeys[idx].items[0]
				monkeys[idx].items = monkeys[idx].items[1:]
				monkeys[idx].inspections++

				value := currentItemWorry
				if monkeys[idx].operation[2] != "old" {
					value = util.ToInt(monkeys[idx].operation[2])
				}
				if monkeys[idx].operation[1] == "+" {
					currentItemWorry += value
				} else {
					currentItemWorry *= value
				}

				newOwner := monkeys[idx].ifFalse
				if currentItemWorry%monkeys[idx].test == 0 {
					newOwner = monkeys[idx].ifTrue
				}

				monkeys[newOwner].items = append(monkeys[newOwner].items, currentItemWorry%upperBound)
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections >= monkeys[j].inspections
	})

	return monkeys[0].inspections * monkeys[1].inspections
}
