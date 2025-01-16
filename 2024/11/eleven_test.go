package eleven

import "testing"

func TestEleven(t *testing.T) {
	if Eleven() != 220357186726677 {
		t.Error("unexpected result")
	}
}
