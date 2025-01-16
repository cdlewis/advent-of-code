package nine

import "testing"

func TestNine(t *testing.T) {
	if Nine(true, `R 5
	U 8
	L 8
	D 3
	R 17
	D 10
	L 25
	U 20`) != 35 {
		t.Fatal("Unexpected result")
	}

	if Nine(false, "") != 2651 {
		t.Fatal("Unexpected result")
	}

}
