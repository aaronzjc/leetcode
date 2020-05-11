package medium

import "testing"

func TestContainerMostWater(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expect := 49

	result := maxArea(input)

	if result != expect {
		t.Fatalf("failed !")
	}

	t.Log("ContainerMostWater passed !")

}