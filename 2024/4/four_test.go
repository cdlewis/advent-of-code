package four

import "testing"

func TestFour(t *testing.T) {
	if Four() != 1972 {
		t.Error("Incorrect output from exercise")
	}	
}