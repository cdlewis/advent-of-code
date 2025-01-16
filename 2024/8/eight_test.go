package eight

import "testing"

func TestEight(t *testing.T) {
	if Eight() != 991 {
		t.Error("expected 991")
	}
}
