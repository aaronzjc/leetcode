package easy

// 地址: https://leetcode-cn.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		minus := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == minus {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
