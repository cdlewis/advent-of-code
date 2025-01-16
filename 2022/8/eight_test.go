package eight

import "testing"

func TestEight(t *testing.T) {
	if Eight(false, "") != 201684 {
		t.Errorf("Unexpected output")
	}

	if Eight(true, `30373
	25512
	65332
	33549
	35390`) != 8 {
		t.Errorf("Unexpected output")
	}
}
