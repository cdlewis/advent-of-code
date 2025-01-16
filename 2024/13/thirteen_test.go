package thirteen

import "testing"

func TestThirteen(t *testing.T) {
	if Thirteen() != 82510994362072 {
		t.Error("unexpected result")
	}
}
