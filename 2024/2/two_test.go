package two

import "testing"

func TestTwo(t *testing.T) {
	if Two() != 717 {
		t.Error("Incorrect output from exercise")
	}
}
