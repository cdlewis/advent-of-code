package eighteen

import (
	"testing"

	"github.com/cdlewis/advent-of-code/util"
)

func TestEighteen(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected int
	}{{
		Input: util.GetInput(18, true, `2,2,2
		1,2,2
		3,2,2
		2,1,2
		2,3,2
		2,2,1
		2,2,3
		2,2,4
		2,2,6
		1,2,5
		3,2,5
		2,1,5
		2,3,5`),
		Expected: 58,
	}, {
		Input:    util.GetInput(18, false, ""),
		Expected: 2074,
	}}

	for _, testCase := range testCases {
		if result := Eighteen(testCase.Input); result != testCase.Expected {
			t.Errorf("Unexpected result: %v", result)
		}
	}
}
