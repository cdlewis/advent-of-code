package twelve

import "testing"

func TestTwelve(t *testing.T) {
	if Twelve() != 862486 {
		t.Error("unexpected result")
	}
}
