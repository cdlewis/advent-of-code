package two

import (
	"iter"
	"slices"
	"strings"

	"github.com/cdlewis/advent-of-code/2024/util"
	"github.com/cdlewis/advent-of-code/2024/util/aoc"
	"github.com/cdlewis/advent-of-code/2024/util/cast"
)

var testData = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func Two() int {
	reports := strings.Split(aoc.GetInput(2, false, testData), "\n")

	safe := 0
	for _, r := range reports {
		report := util.Map(strings.Split(r, " "), cast.ToInt[string])

		valid := checkReport(report)
		if valid {
			safe++
			continue
		}

		slices.Reverse(report)
		reverseValid := checkReport(report)
		if reverseValid {
			safe++
		}
	}

	return safe
}

func checkReport(reports []int) bool {
CHECK_REPORTS:
	for r := range reportPermutations(reports) {
		for i, j := 0, 1; j < len(r); i, j = j, j+1 {
			if r[i] == r[j] || r[j]-r[i] < 0 || r[j]-r[i] > 3 {
				continue CHECK_REPORTS
			}
		}

		return true
	}

	return false
}

func reportPermutations(report []int) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		for i := -1; i < len(report); i++ {
			newReport := report
			if i > -1 {
				newReport = slices.Concat(newReport[:i], newReport[i+1:])
			}

			if !yield(newReport) {
				return
			}
		}
	}
}
