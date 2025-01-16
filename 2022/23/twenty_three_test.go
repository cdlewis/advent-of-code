package twenty_three

import "testing"

func TestTwentyThree(t *testing.T) {
	if result := TwentyThree(); result != 893 {
		t.Errorf("Unexpected result: %v", result)
	}
}
