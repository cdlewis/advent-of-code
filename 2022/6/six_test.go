package six

import "testing"

func TestSix(t *testing.T) {
	if Six(false, "") != 2472 {
		t.Error("Incorrect output from exercise")
	}

	if Six(true, "mjqjpqmgbljsphdztnvjfqwrcgsmlb") != 19 {
		t.Error("Incorrect output from exercise")
	}
}
