package easy

// https://leetcode-cn.com/problems/roman-to-integer/

func romanToInt(s string) (res int) {
	m := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	var c, n string
	i, l, rs := 0, len(s)-1, []byte(s)
	for i <= l {
		if n != "" {
			c = n
		} else {
			c = string(rs[i])
		}
		i++
		if i > l {
			res += m[c]
			continue
		}
		n = string(rs[i])
		if r, ok := m[c+n]; ok {
			res += r
			n = ""
			i++
			continue
		}
		res += m[c]
	}
	return
}
