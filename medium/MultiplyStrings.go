package medium

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/multiply-strings/

// TIPS:
// 按照乘法的运算规则进行计算。思路大同小异。优化空间在于，字符串转化。

func multiply(num1 string, num2 string) (res string) {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	// 首先将字符串转化成整数数组
	var a, b []int
	as, bs := strings.Split(num1, ""), strings.Split(num2, "")
	var i, e int
	for {
		if i < len(as) {
			e, _ = strconv.Atoi(as[i])
			a = append(a, e)
		}
		if i < len(bs) {
			e, _ = strconv.Atoi(bs[i])
			b = append(b, e)
		}
		if i >= len(as)-1 && i >= len(bs)-1 {
			break
		}
		i++
	}

	var pass, val, m int
	var cache [][]int

	for i := len(a) - 1; i >= 0; i-- {
		cache2 := []int{}
		for j := len(b) - 1; j >= 0; j-- {
			m = a[i]*b[j] + pass
			pass, val = m/10, m%10
			cache2 = append(cache2, val)
		}
		if pass > 0 {
			cache2 = append(cache2, pass)
		}
		pad := make([]int, len(a)-1-i)
		cache2 = append(pad, cache2...)
		cache = append(cache, cache2)
		pass, val, m = 0, 0, 0
	}

	var resInt []int

	var end bool
	pass, val, m = 0, 0, 0
	j := 0
	for {
		end = true
		for i := 0; i < len(cache); i++ {
			if j < len(cache[i]) {
				end = false
				m += cache[i][j]
			}
		}
		if end {
			break
		}
		m += pass
		val, pass = m%10, m/10
		resInt = append(resInt, val)
		m = 0
		j++
	}
	if pass > 0 {
		resInt = append(resInt, pass)
	}
	for i := len(resInt) - 1; i >= 0; i-- {
		res += strconv.Itoa(resInt[i])
	}
	return
}
