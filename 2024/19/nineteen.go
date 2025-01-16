package main

import (
	"fmt"
	"strings"

	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/set"
)

func main() {
	raw := strings.Split(aoc.GetInput(19, false, ""), "\n\n")
	have := set.New(strings.Split(raw[0], ", ")...)
	want := strings.Split(raw[1], "\n")

	results := 0
	for _, w := range want {
		results += canAssemble(w, have)
	}

	fmt.Println(results)
}

var cache = map[string]int{}

func canAssemble(want string, have set.Set[string]) int {
	if want == "" {
		return 1
	}

	if result, ok := cache[want]; ok {
		return result
	}

	perms := 0

	for i := 0; i < len(want)+1; i++ {
		if have.Exists(want[:i]) {
			perms += canAssemble(want[i:], have)
		}
	}

	cache[want] = perms
	return perms
}
