package leetcode

// Using Backtracking with Bitmask
func countArrangementBitMask(n int) int {
	count := 0

	var dfs func(int, int)
	dfs = func(mask int, pos int) {
		if pos > n {
			count++
		}
		for i := 1; i <= n; i++ {
			if mask&(1<<i) == 0 && (pos%i == 0 || i%pos == 0) {
				dfs(mask|(1<<i), pos+1)
			}
		}
	}

	dfs(0, 1)
	return count
}

// Time: O(k) k numbers of valid permutations
func countArrangement(n int) int {
	A := make([]int, n)
	for i := 1; i <= n; i++ {
		A[i-1] = i
	}
	count := 0
	permute(A, 0, &count)
	return count
}

func permute(A []int, p int, count *int) {
	if p == len(A) {
		(*count)++
	}
	for i := p; i < len(A); i++ {
		A[i], A[p] = A[p], A[i]
		if A[p]%(p+1) == 0 || (p+1)%A[p] == 0 {
			permute(A, p+1, count)
		}
		A[i], A[p] = A[p], A[i]
	}
}
