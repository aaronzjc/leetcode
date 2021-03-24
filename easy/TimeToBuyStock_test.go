package easy

import "testing"

func TestMaxProfit(t *testing.T) {
	seeds := []struct {
		input  []int
		expect int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, v := range seeds {
		result := maxProfit(v.input)
		if result != v.expect {
			t.Error(v, result)
			t.Fatal("failed !")
		}
	}

	t.Log("TimeToBuyStock passed !")
}
