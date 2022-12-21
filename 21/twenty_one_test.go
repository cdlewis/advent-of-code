package TwentyOne

import "testing"

func TestTwentyOne(t *testing.T) {
	if TwentyOne() != 3617613952378 {
		t.Errorf("Unexpected result")
	}
}
