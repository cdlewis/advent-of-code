package twelve

import (
	"testing"

	"github.com/cdlewis/advent-of-code/util"
)

func TestTwelve(t *testing.T) {
	if result := Twelve(false, ""); result != 500 {
		t.Errorf("Unexpected output: " + util.ToString(result))
	}

	if result := Twelve(true, `Sabqponm
	abcryxxl
	accszExk
	acctuvwj
	abdefghi`); result != 29 {
		t.Errorf("Unexpected output: " + util.ToString(result))
	}
}
