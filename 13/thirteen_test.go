package thirteen

import "testing"

func TestThirteen(t *testing.T) {
	if result := Thirteen(false, ""); result != 20304 {
		t.Errorf("Unexpected output: %v", result)
	}

	if result := Thirteen(true, `[1,1,3,1,1]
	[1,1,5,1,1]
	
	[[1],[2,3,4]]
	[[1],4]
	
	[9]
	[[8,7,6]]
	
	[[4,4],4,4]
	[[4,4],4,4,4]
	
	[7,7,7,7]
	[7,7,7]
	
	[]
	[3]
	
	[[[]]]
	[[]]
	
	[1,[2,[3,[4,[5,6,7]]]],8,9]
	[1,[2,[3,[4,[5,6,0]]]],8,9]`); result != 140 {
		t.Errorf("Unexpected output: %v", result)
	}
}
