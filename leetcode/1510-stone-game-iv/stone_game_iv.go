package leetcode

func winnerSquareGame(n int) bool {
	D := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if !D[i-j*j] {
				D[i] = true
				break
			}
		}
	}
	return D[n]
}
