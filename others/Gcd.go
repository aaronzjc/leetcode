package others

// TIPS:
// 根据辗转相除法: gcd(a, b) = gcd (b, a % b)

func gcd(a int, b int) int {
	var r int
	if a < b {
		a, b = b, a
	}
	for {
		r = a % b
		if r == 0 {
			break
		}
		a, b = b, r
	}

	return b
}
