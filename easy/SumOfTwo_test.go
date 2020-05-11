package easy

import (
	"reflect"
	"testing"
)

func TestSumTwo(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expect := []int{0, 1}

	result := twoSum(nums, target)

	if !reflect.DeepEqual(result, expect) {
		t.Error("failed !")
	}
	t.Log("SumTwo passed")
}
