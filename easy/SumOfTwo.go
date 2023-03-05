package easy

// https://leetcode-cn.com/problems/two-sum/

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

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if idx, ok := m[target-nums[i]]; ok {
			return []int{i, idx}
		}
		m[nums[i]] = i
	}
	return []int{}
}
