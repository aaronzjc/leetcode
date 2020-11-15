package easy

// https://leetcode-cn.com/problems/plus-one/

func plusOne(digits []int) []int {
	add := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + add
		if digits[i] >= 10 {
			digits[i] = digits[i] - 10
			add = 1
		} else {
			add = 0
		}
	}
	if add > 0 {
		digits = append(digits, 0)
		copy(digits[1:], digits[:])
		digits[0] = 1
	}

	return digits
}
