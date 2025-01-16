package five

import "testing"

func TestFive(t *testing.T) {
	if Five() != 6897 {
		t.Error("Incorrect output from exercise")
	}
}
