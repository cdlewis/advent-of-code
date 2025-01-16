package nineteen

import "testing"

func TestNinteen(t *testing.T) {
	if Nineteen() != 4114 {
		t.Errorf("Unexpected output")
	}
}
