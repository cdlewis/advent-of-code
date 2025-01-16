package ten

import "testing"

func TestTen(t *testing.T) {
	if Ten() != 1225 {
		t.Error("unexpected result")
	}
}
