package three

import "testing"

func TestThree(t *testing.T) {
	if Three() != 2650 {
		t.Error("Incorrect output from exercise")
	}
}
