package leetcode

const maxInt = 1<<31 - 1

// O(n*k^2)
func minCostII(costs [][]int) int {
	n := len(costs)
	if n == 0 {
		return 0
	}
	k := len(costs[0])
	D, T := make([]int, k), make([]int, k)
	copy(D, costs[0])
	for i := 1; i < n; i++ {
		for j := 0; j < k; j++ {
			T[j] = maxInt
			for h := 0; h < k; h++ {
				if j != h {
					T[j] = min(T[j], costs[i][j]+D[h])
				}
			}
		}
		copy(D, T)
	}

	answer := D[0]
	for j := 1; j < k; j++ {
		answer = min(answer, D[j])
	}
	return answer
}

// O(n*k)
func minCostIIOptimized(costs [][]int) int {
	n := len(costs)
	if n == 0 {
		return 0
	}
	k := len(costs[0])

	D := make([]int, k)
	min1, min2 := 0, 0
	for i := 0; i < n; i++ {
		oldMin1, oldMin2 := min1, min2
		min1 = maxInt
		min2 = maxInt
		for j := 0; j < k; j++ {
			if D[j] != oldMin1 || oldMin1 == oldMin2 {
				D[j] = oldMin1 + costs[i][j]
			} else {
				D[j] = oldMin2 + costs[i][j]
			}
			if min1 <= D[j] {
				min2 = min(min2, D[j])
			} else {
				min2 = min1
				min1 = D[j]
			}
		}
	}
	return min1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
