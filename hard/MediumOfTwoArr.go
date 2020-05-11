package hard

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	var small, large []int
	var dst []float64
	if len1 <= len2 {
		small, large = nums1, nums2
	} else {
		small, large = nums2, nums1
	}
	var left, right int
	if (len1 + len2) % 2 != 0 {
		left = (len1 + len2) / 2
		right = left
	} else {
		left = (len1 + len2 - 1) / 2
		right = (len1 + len2 + 1) / 2
	}
	j := 0
	for i := 0; i < len(large);i++ {
		for j < len(small) {
			if small[j] > large[i] {
				break
			}
			dst = append(dst, float64(small[j]))
			j++
		}
		dst = append(dst, float64(large[i]))
	}
	if j < len(small) {
		dst = append(dst, float64(small[j]))
		j++
	}
	return (dst[left] + dst[right]) / 2
}