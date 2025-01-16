package six

import (
	"fmt"
	"testing"
)

func TestSix(t *testing.T) {
	if result := Six(); result != 1424 {
		t.Error(fmt.Sprintf("Incorrect output from exercise: %v", result))
	}
}
