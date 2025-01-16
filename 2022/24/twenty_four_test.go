package twenty_four

import "testing"

func TestTwentyFour(t *testing.T) {
	if result := TwentyFour(); result != 728 {
		t.Errorf("Unexpected result: %v", result)
	}
}
