package hard

// https://leetcode-cn.com/problems/insert-interval/

// TIPS:
// 比中等难度的合并区间更加简单，因为这个排序好了。这个题思路清晰，代码清晰，赞一个。

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	i := len(intervals) - 1
	for i > 0 {
		cur := intervals[i]
		pre := intervals[i-1]
		// 在右边
		if cur[0] > pre[1] {
			break
		}
		// 在左边
		if cur[0] < pre[0] && cur[1] < pre[0] {
			intervals[i-1], intervals[i] = intervals[i], intervals[i-1]
			i--
			continue
		}

		var tmp []int
		// 右交叉
		if cur[0] >= pre[0] && cur[1] >= pre[1] {
			tmp = []int{pre[0], cur[1]}
			intervals[i] = tmp
		}
		// 左交叉
		if cur[0] <= pre[0] && cur[1] <= pre[1] {
			tmp = []int{cur[0], pre[1]}
			intervals[i] = tmp
		}
		// 在中间
		if cur[0] > pre[0] && cur[1] < pre[1] {
			intervals[i-1], intervals[i] = intervals[i], intervals[i-1]
		}
		copy(intervals[i-1:], intervals[i:])
		intervals = intervals[:len(intervals)-1]
		i--
	}

	return intervals
}
