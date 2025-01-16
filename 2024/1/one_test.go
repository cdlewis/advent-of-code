package one

import "testing"

func TestOne(t *testing.T) {
	if One() != 20351745 {
		t.Error("Incorrect output from exercise")
	}
}
