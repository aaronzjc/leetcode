package tools

import (
	"encoding/json"
	"reflect"
	"sort"
)

// IsIntArrEquals 比较两个整形数组是否相等
func IsIntArrEquals(a []int, b []int, order bool) bool {
	if len(a)+len(b) == 0 {
		return true
	}
	if !order {
		sort.Ints(a)
		sort.Ints(b)
	}
	return reflect.DeepEqual(a, b)
}

// IsStringArrEquals 比较两个字符串数组是否相等
func IsStringArrEquals(a []string, b []string, order bool) bool {
	if len(a)+len(b) == 0 {
		return true
	}
	if !order {
		sort.Strings(a)
		sort.Strings(b)
	}
	return reflect.DeepEqual(a, b)
}

// 采用反射实现的一个比较，比较复杂，也不建议使用
//func isArrOrSliceEquals(a interface{}, b interface{}, order bool) bool {
//	var t int
//	ta, tb := reflect.TypeOf(a).Kind(), reflect.TypeOf(b).Kind()
//	va, vb := reflect.ValueOf(a), reflect.ValueOf(b)
//	if ta != tb {
//		return false
//	}
//	if va.Len() != vb.Len() {
//		return false
//	}
//	if va.Len() == vb.Len() && va.Len() == 0 {
//		return true
//	}
//	if ta != reflect.Slice && ta != reflect.Array {
//		// 不支持的类型比较
//		return false
//	}
//	switch reflect.TypeOf(va.Index(0)).Kind() {
//	case reflect.Int:
//		t = 0
//	case reflect.String:
//		t = 1
//	}
//
//	var i int
//	var ia, ib []int
//	var sa, sb []string
//	for i < va.Len() || i < vb.Len() {
//		if i < va.Len() {
//			if t == 0 {
//				ia = append(ia, int(va.Index(i).Int()))
//			} else if t == 1 {
//				sa = append(sa, va.Index(i).String())
//			}
//		}
//		if i < vb.Len() {
//			if t == 0 {
//				ib = append(ib, int(va.Index(i).Int()))
//			} else if t == 1 {
//				sb = append(sb, va.Index(i).String())
//			}
//		}
//		i++
//	}
//	// 如果不比较顺序的化，这里先进行排序
//	if t == 0 {
//		if !order {
//			sort.Ints(ia)
//			sort.Ints(ib)
//		}
//		return reflect.DeepEqual(ia, ib)
//	} else if t == 1 {
//		if !order {
//			sort.Strings(sa)
//			sort.Strings(sb)
//		}
//		return reflect.DeepEqual(sa, sb)
//	}
//	return false
//}

// IsL2IntArrayEquals 比较两个整形二维数组是否相等
func IsL2IntArrayEquals(a [][]int, b [][]int, order bool) bool {
	var flatA, flatB []string
	var i int
	for i < len(a) || i < len(b) {
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
		i++
	}
	return IsStringArrEquals(flatA, flatB, order)
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

	return IsStringArrEquals(flatA, flatB, order)
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
