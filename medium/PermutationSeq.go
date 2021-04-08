package medium

import (
	"strconv"
)

// https://leetcode-cn.com/problems/permutation-sequence/

// TIPS:
// 回溯法

func getPermutation(n int, k int) (res string) {
	var counter int

	var permute func([]int, string)
	permute = func(cc []int, str string) {
		if res != "" {
			return
		}
		if len(str) == n {
			counter++
			if counter == k {
				res = str
			}
		}
		for k, v := range cc {
			tmp := make([]int, len(cc))
			copy(tmp, cc)
			copy(tmp[k:], tmp[k+1:])
			tmp = tmp[:len(tmp)-1]
			permute(tmp, str+strconv.Itoa(v))
		}
	}
	var cc []int
	for i := 1; i <= n; i++ {
		cc = append(cc, i)
	}
	permute(cc, "")

	return
}
