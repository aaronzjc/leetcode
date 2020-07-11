package medium

// https://leetcode-cn.com/problems/merge-intervals/

// TIPS:
// 苦力活。每次选择一个元素，放到结果集中。然后对结果集，从右往左，两两合并。有点像数组选择排序的思路。

type mstruct struct {
	left  int
	right int
}

func merge(intervals [][]int) (res [][]int) {
	var resMap []mstruct

	for _, v := range intervals {
		ele := mstruct{
			left:  v[0],
			right: v[1],
		}
		resMap = append(resMap, ele)

		var i int
		i = len(resMap) - 1
		for i > 0 {
			cur := resMap[i]
			pre := resMap[i-1]
			// 在右边
			if cur.left > pre.right {
				break
			}
			// 在左边
			if cur.left < pre.left && cur.right < pre.left {
				resMap[i-1], resMap[i] = resMap[i], resMap[i-1]
				i--
				continue
			}

			var tmp mstruct
			// 右交叉
			if cur.left >= pre.left && cur.right >= pre.right {
				tmp = mstruct{
					left:  pre.left,
					right: cur.right,
				}
				resMap[i] = tmp
			}
			// 左交叉
			if cur.left <= pre.left && cur.right <= pre.right {
				tmp = mstruct{
					left:  cur.left,
					right: pre.right,
				}
				resMap[i] = tmp
			}
			// 在中间
			if cur.left > pre.left && cur.right < pre.right {
				resMap[i-1], resMap[i] = resMap[i], resMap[i-1]
			}
			copy(resMap[i-1:], resMap[i:])
			resMap = resMap[:len(resMap)-1]
			i--
		}
	}
	for _, v := range resMap {
		res = append(res, []int{v.left, v.right})
	}
	return
}
