package leetcode

func minCostDP(costs [][]int) int {
	n := len(costs)
	if n == 0 {
		return 0
	}
	D := make([]int, 3)
	copy(D, costs[0])
	for i := 1; i < n; i++ {
		T := []int{0, 0, 0}
		for j := 0; j < 3; j++ {
			T[j] = costs[i][j] + min(D[(j+1)%3], D[(j+2)%3])
		}
		D = T
	}
	return min(D[0], min(D[1], D[2]))
}

func minCost(costs [][]int) int {
	n := len(costs)
	if n == 0 {
		return 0
	}
	r, b, g := 0, 0, 0
	for i := 0; i < n; i++ {
		rc, bc, gc := costs[i][0], costs[i][1], costs[i][2]
		rc += min(b, g)
		bc += min(r, g)
		gc += min(r, b)
		r, b, g = rc, bc, gc
	}
	return min(r, min(b, g))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
