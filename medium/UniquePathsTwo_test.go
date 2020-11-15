package medium

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	seeds := []struct {
		grid   [][]int
		expect int
	}{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
		{[][]int{{0, 1}, {0, 0}}, 1},
		{[][]int{{0, 0}}, 1},
	}

	for _, v := range seeds {
		result := uniquePathsWithObstacles(v.grid)
		if result != v.expect {
			t.Error(v, result)
			t.Fatalf("failed !")
		}
	}

	t.Log("UniquePathsTwo passed !")
}
