package leetcode

func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	if m == 0 {
		return n == 0
	}

	D := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		D[i] = make([]bool, m+1)
	}

	D[0][0] = true
	for j := 1; j <= m; j++ {
		if p[j-1] == '*' {
			D[0][j] = D[0][j-1]
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			switch p[j-1] {
			case '*':
				D[i][j] = D[i][j-1] || D[i-1][j]
			case '?', s[i-1]:
				D[i][j] = D[i-1][j-1]
			default:
				D[i][j] = false
			}
		}
	}

	return D[n][m]
}
