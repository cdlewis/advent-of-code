package twenty

import "testing"

func TestTwenty(t *testing.T) {
	if Twenty() != 831878881825 {
		t.Errorf("Unexpected result")
	}
}
