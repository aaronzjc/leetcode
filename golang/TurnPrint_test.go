package golang

import "testing"

func TestTurnPrint(t *testing.T) {
	seeds := []struct {
		input  int
		expect string
	}{
		{1, "1A"},
		{2, "1A2B"},
		{4, "1A2B3C4D"},
		{10, "1A2B3C4D5E6F7G8H9I1J"},
		{27, "1A2B3C4D5E6F7G8H9I1J2K3L4M5N6O7P8Q9R1S2T3U4V5W6X7Y8Z9A"},
	}
	for _, v := range seeds {
		result := TurnPrint(v.input)
		if v.expect != result {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}
	t.Log("TurnPrint passed !")
}
