package easy

import (
	"strconv"
)

// https://leetcode-cn.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	sx := strconv.Itoa(x)
	i, j := 0, len(sx)-1
	for i <= j {
		if sx[i] != sx[j] {
			return false
		}
		i++
		j--
	}
	return true
}
