package five

import "testing"

func TestFive(t *testing.T) {
	if Five() != "GSLCMFBRP" {
		t.Error("Incorrect output from exercise")
	}
}
