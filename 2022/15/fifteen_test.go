package fifteen

import "testing"

func TestFifteen(t *testing.T) {
	if result := Fifteen(); result != 11557863040754 {
		t.Errorf("Unexpected result %v", result)
	}
}
