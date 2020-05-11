package medium

// https://leetcode-cn.com/problems/string-to-integer-atoi/

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	bs := strings.TrimSpace(str)

	prefix := 1
	result := 0

	bk := false

	for i := 0;i < len(bs);i++ {
		switch bs[i] {
		case 43:
			if i != 0 {
				bk = true
			}
			break
		case 45:
			if i == 0 {
				prefix = -1
			} else {
				bk = true
			}
			break
		case 48,49,50,51,52,53,54,55,56,57:
			result = result * 10 + (int(bs[i]) - '0')
			if result*prefix > math.MaxInt32 {
				return math.MaxInt32
			}
			if result*prefix < ^math.MaxInt32 {
				return ^math.MaxInt32
			}
			break
		default:
			bk = true
			break
		}
		if bk {
			break
		}
	}
	return result * prefix
}