package medium

import "testing"

func TestCoinChange(t *testing.T) {
	seeds := []struct {
		coins  []int
		amount int
		expect int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{1}, 1, 1},
		{[]int{1}, 2, 2},
	}

	for _, v := range seeds {
		result := CoinChange(v.coins, v.amount)
		if result != v.expect {
			t.Error(v, result)
			t.Fatal("failed !")
		}
	}

	t.Log("CoinChange passed !")
}
