package tools

import (
	"encoding/json"
	"reflect"
	"sort"
)

type AnySliceT interface {
	int | float64 | string
}

// IsArrEquals 比较两个数组是否相等
func IsArrEquals[T AnySliceT](a []T, b []T, order bool) bool {
	if len(a)+len(b) == 0 {
		return true
	}
	if !order {
		sort.Slice(a, func(i, j int) bool {
			return a[i] >= a[j]
		})
		sort.Slice(b, func(i, j int) bool {
			return b[i] >= b[j]
		})
	}
	return reflect.DeepEqual(a, b)
}

// IsL2IntArrayEquals 比较两个整形二维数组是否相等
func IsL2IntArrayEquals(a [][]int, b [][]int, order bool) bool {
	var flatA, flatB []string
	if len(a) != len(b) {
		return false
	}
	if len(a)+len(b) == 0 {
		return true
	}
	for i := 0; i < len(a); i++ {
		if i < len(a) {
			if !order {
				sort.Ints(a[i])
			}
			bs, _ := json.Marshal(a[i])
			flatA = append(flatA, string(bs))
		}
		if i < len(b) {
			if !order {
				sort.Ints(b[i])
			}
			bs, _ := json.Marshal(b[i])
			flatB = append(flatB, string(bs))
		}
	}
	return IsArrEquals(flatA, flatB, order)
}

// IsL2ByteArrayEquals 比较两个二维字节数组是否相等
func IsL2ByteArrayEquals(aByte [][]byte, bByte [][]byte, order bool) bool {
	var aInt, bInt [][]int
	aInt = make([][]int, len(aByte))
	for i := 0; i < len(aByte); i++ {
		aInt[i] = make([]int, len(aByte[i]))
		for j := 0; j < len(aByte[i]); j++ {
			aInt[i][j] = int(aByte[i][j])
		}
	}
	bInt = make([][]int, len(bByte))
	for i := 0; i < len(bByte); i++ {
		bInt[i] = make([]int, len(bByte[i]))
		for j := 0; j < len(bByte[i]); j++ {
			bInt[i][j] = int(bByte[i][j])
		}
	}
	return IsL2IntArrayEquals(aInt, bInt, order)
}

// IsL2StrArrayEquals 比较两个二维字符串数组是否相等
func IsL2StrArrayEquals(aStr [][]string, bStr [][]string, order bool) bool {
	var flatA, flatB []string
	var i int
	for i < len(aStr) || i < len(bStr) {
		if i < len(aStr) {
			if !order {
				sort.Strings(aStr[i])
			}
			bs, _ := json.Marshal(aStr[i])
			flatA = append(flatA, string(bs))
		}
		if i < len(bStr) {
			if !order {
				sort.Strings(bStr[i])
			}
			bs, _ := json.Marshal(bStr[i])
			flatB = append(flatB, string(bs))
		}
		i++
	}

	return IsArrEquals(flatA, flatB, order)
}

// Permutation 单词的全排列
func Permutation(words []string) []string {
	var ls []string
	var per func([]string, string)
	per = func(words []string, res string) {
		if len(words) == 0 {
			if res != "" {
				ls = append(ls, res)
			}
			return
		}

		for i := 0; i < len(words); i++ {
			var tmp []string
			if i > 0 {
				tmp = append(tmp, words[:i]...)
			}
			if i < len(words)-1 {
				tmp = append(tmp, words[i+1:]...)
			}
			per(tmp, res+words[i])
		}
	}

	per(words, "")

	return ls
}
