package five

import (
	"slices"
	"strings"

	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/cast"
)

func Five() int {
	rawInput := strings.Split(aoc.GetInput(5, false, ""), "\n")
	sectionEnd := slices.Index(rawInput, "")
	rawRules := rawInput[:sectionEnd]
	rawManuals := rawInput[sectionEnd+1:]

	deps := map[string][]string{}
	for _, r := range rawRules {
		tokens := strings.Split(r, "|")
		deps[tokens[1]] = append(deps[tokens[1]], tokens[0])
	}

	manuals := make([][]string, 0, len(rawManuals))
	for _, m := range rawManuals {
		manuals = append(manuals, strings.Split(m, ","))
	}

	sumMid := 0

	for _, manual := range manuals {
		seen := map[string]int{}
		for idx, page := range manual {
			seen[page] = idx
		}

		valid := true
	CHECK_PAGES:
		for pageIdx, page := range manual {
			for _, requiredPage := range deps[page] {
				if seenIdx, ok := seen[requiredPage]; ok && seenIdx > pageIdx {
					valid = false
					break CHECK_PAGES
				}
			}
		}

		// we don't care about correctly ordered manuals
		if valid {
			continue
		}

		newManual := make([]string, 0, len(manual))
		usedPage := make(map[string]struct{}, len(manual))
		var stack []string
		for _, page := range manual {
			hasDeps := false
			for _, requiredPage := range deps[page] {
				if _, ok := seen[requiredPage]; ok {
					hasDeps = true
					break
				}
			}

			if hasDeps {
				stack = append(stack, page)
				continue
			}

			usedPage[page] = struct{}{}
			newManual = append(newManual, page)
		}

		for len(stack) > 0 {
			curr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if _, ok := usedPage[curr]; ok {
				continue
			}

			readyToPlace, needsPages := dependenciesSatisfied(
				curr,
				deps,
				seen,
				usedPage,
			)
			if readyToPlace {
				newManual = append(newManual, curr)
				usedPage[curr] = struct{}{}
				continue
			}

			stack = append(stack, curr)
			stack = append(stack, needsPages...)
		}

		sumMid += cast.ToInt(newManual[len(newManual)/2])
	}

	return sumMid
}

func dependenciesSatisfied(
	page string,
	dependencies map[string][]string,
	present map[string]int,
	seen map[string]struct{},
) (bool, []string) {
	var remainingDependencies []string

	for _, requiredPage := range dependencies[page] {
		if _, ok := present[requiredPage]; !ok {
			continue
		}

		if _, ok := seen[requiredPage]; ok {
			continue
		}

		remainingDependencies = append(remainingDependencies, requiredPage)
	}

	return len(remainingDependencies) == 0, remainingDependencies
}
