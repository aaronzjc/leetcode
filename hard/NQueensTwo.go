package hard

// https://leetcode-cn.com/problems/n-queens/

// TIPS:
// 和上一题一模一样呀，顺手刷了！
// 但是我的性能比其他的回溯法要差，没仔细研究其他人的方法。日后回过头再分析吧。

func totalNQueens(n int) int {
	var solve [][][]int
	var find func(int, int, [][]int)
	find = func(x int, y int, res [][]int) {
		if len(res) == n {
			solve = append(solve, res)
			return
		}

		var i, j int
		j = y
		for i = x; i < n; i++ {
			j = j % n
			for ; j < n; j++ {
				if i == -1 || j == -1 {
					continue
				}
				ur := true
				for k := len(res) - 1; k >= 0; k-- {
					ii := i - res[k][0]
					jj := j - res[k][1]
					if jj == 0 || ii == 0 || ii*ii == jj*jj {
						ur = false
					}
					if !ur {
						break
					}
				}
				if !ur {
					continue
				}
				tmp := make([][]int, len(res))
				copy(tmp, res)
				tmp = append(tmp, []int{i, j})
				find(i, j, tmp)
			}
		}
	}
	find(-1, -1, nil)
	return len(solve)
}
