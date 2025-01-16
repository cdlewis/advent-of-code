package four

import "testing"

func TestFour(t *testing.T) {
	if Four() != 905 {
		t.Error("Incorrect output from exercise")
	}
}
