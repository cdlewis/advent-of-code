package twenty_five

import "testing"

func TestTwentyFive(t *testing.T) {
	if result := TwentyFive(); result != "2=--=0000-1-0-=1=0=2" {
		t.Errorf("Unexpected result: %v", result)
	}
}
