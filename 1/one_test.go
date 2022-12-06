package one

import "testing"

func TestOne(t *testing.T) {
	if One() != 205615 {
		t.Error("Incorrect output from exercise")
	}
}
