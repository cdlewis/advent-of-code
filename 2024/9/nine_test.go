package nine

import "testing"

func TestNine(t *testing.T) {
	if Nine() != 6360363199987 {
		t.Error("incorrect result")
	}
}
