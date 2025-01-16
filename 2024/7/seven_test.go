package seven

import "testing"

func TestSeven(t *testing.T) {
	if Seven() != 354060705047464 {
		t.Error("Invalid result")
	}
}
