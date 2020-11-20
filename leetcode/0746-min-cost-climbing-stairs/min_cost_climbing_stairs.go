package leetcode

func minCostClimbingStairsDP(cost []int) int {
	n := len(cost)
	if n <= 1 {
		return 0
	}
	D := make([]int, n)
	D[0] = cost[0]
	D[1] = cost[1]
	for i := 2; i < n; i++ {
		D[i] = cost[i] + min(D[i-1], D[i-2])
	}
	return min(D[n-1], D[n-2])
}

// f(i) = cost(i) + min(f(i+1), f(i+2))
func minCostClimbingStairs(cost []int) int {
	f1, f2 := 0, 0
	for i := len(cost) - 1; i >= 0; i-- {
		f := cost[i] + min(f1, f2)
		f2 = f1
		f1 = f
	}
	return min(f1, f2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
