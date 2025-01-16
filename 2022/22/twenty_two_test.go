package twenty_two

import "testing"

func TestTwentyTwo(t *testing.T) {
	if result := TwentyTwo(); result != 11451 {
		t.Errorf("Unexpected result: %v", result)
	}
}
