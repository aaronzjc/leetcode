package hard

// https://leetcode-cn.com/problems/n-queens/

// TIPS:
// 回溯法。回溯法的思路是最清晰的，遇到走一步看一步时，一般都是回溯法。

func solveNQueens(n int) [][]string {
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
	res := make([][]string, len(solve))
	for i := 0; i < len(solve); i++ {
		res[i] = make([]string, n)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				char := "."
				for s := 0; s < len(solve[i]); s++ {
					if solve[i][s][0] == j && solve[i][s][1] == k {
						char = "Q"
						break
					}
				}
				res[i][j] += char
			}
		}
	}

	return res
}
