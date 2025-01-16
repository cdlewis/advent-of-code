package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/cdlewis/advent-of-code/2024/util"
	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/set"
)

func main() {
	input := strings.Split(aoc.GetInput(23, false, ""), "\n")
	pairExists := set.New[string]()
	keys := set.New[string]()
	for _, i := range input {
		tokens := strings.Split(i, "-")
		keys.Add(tokens[0])
		keys.Add(tokens[1])
		pairExists.Add(tokens[0] + tokens[1])
		pairExists.Add(tokens[1] + tokens[0])
	}

	candidates := util.FromIter(maps.Keys(keys))

	result := biggestSet(set.New[string](), candidates, 0, pairExists)
	slices.Sort(result)
	fmt.Println(strings.Join(result, ","))
}

func biggestSet(includeComputers set.Set[string], candidates []string, candidateIndex int, pairs set.Set[string]) []string {
	if candidateIndex >= len(candidates) {
		return util.FromIter(maps.Keys(includeComputers))
	}

	nextComputer := candidates[candidateIndex]
	var bestSetWithNextComputer []string
	canInclude := includeComputers.ForAll(func(c string) bool {
		return pairs.Exists(nextComputer + c)
	})
	if canInclude {
		newIncludeComputers := maps.Clone(includeComputers)
		newIncludeComputers.Add(nextComputer)
		bestSetWithNextComputer = biggestSet(newIncludeComputers, candidates, candidateIndex+1, pairs)
	}

	bestSetWithoutNextComputer := biggestSet(includeComputers, candidates, candidateIndex+1, pairs)

	if len(bestSetWithNextComputer) > len(bestSetWithoutNextComputer) {
		return bestSetWithNextComputer
	}

	return bestSetWithoutNextComputer
}
