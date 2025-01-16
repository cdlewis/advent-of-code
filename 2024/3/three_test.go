package three

import "testing"

func TestThree(t *testing.T) {
	if Three() != 90772405 {
		t.Error("Incorrect output from exercise")
	}	
}