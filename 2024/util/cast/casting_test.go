package cast

import "testing"

func TestToInt(t *testing.T) {
	if ToInt(rune('7')) != 7 {
		t.Errorf("ToInt(rune('7')) expected 7")
	}

	if ToInt(byte('7')) != 7 {
		t.Errorf("ToInt(byte('7')) expected 7")
	}

	if ToInt("676") != 676 {
		t.Errorf("ToInt(byte('676')) expected 676")
	}

	if ToInt([]byte("676")) != 676 {
		t.Errorf("ToInt(byte('676')) expected 676")
	}
}

func TestToString(t *testing.T) {
	byteTests := []struct {
		name  string
		input interface{}
		want  string
	}{
		{"byte", byte('a'), "a"},
		{"byte", byte('x'), "x"},
		{"int", 1234, "1234"},
		{"int", 512, "512"},
		{"rune", rune(65), "A"},
		{"rune", rune(97), "a"},
	}
	for _, tt := range byteTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.input); got != tt.want {
				t.Errorf("ToString(byte) = %q, want %q", got, tt.want)
			}
		})
	}
}
