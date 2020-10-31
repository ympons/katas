package leetcode

var table = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

func romanToInt(s string) int {
	num, n := 0, len(s)
	for i := 0; i < n; {
		x := 0
		if j := i + 2; j <= n {
			x = table[string(s[i:i+2])]
		}
		if x > 0 {
			num += x
			i += 2
		} else {
			num += table[string(s[i])]
			i++
		}
	}

	return num
}
