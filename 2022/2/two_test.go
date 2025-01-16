package two

import "testing"

func TestTwo(t *testing.T) {
	if Two() != 9541 {
		t.Error("Incorrect output from exercise")
	}
}
