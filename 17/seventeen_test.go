package seventeen

import (
	"testing"

	"github.com/cdlewis/advent-of-code/util"
)

func TestSeventeen(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>",
			expected: 1514285714288,
		},
		{
			input:    util.GetInput(17, false, ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"),
			expected: 1597714285698,
		},
	}

	for _, testCase := range testCases {
		if result := Seventeen(testCase.input); result != testCase.expected {
			t.Fatalf("Unexpected answer %v", result)
		}
	}
}
