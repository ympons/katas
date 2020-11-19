package leetcode

func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	D := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		D[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				D[i][j] = D[i-1][j-1] + 1
			} else {
				D[i][j] = D[i][j-1]
				if D[i-1][j] > D[i][j] {
					D[i][j] = D[i-1][j]
				}
			}
		}
	}
	return D[n][m]
}
