package leetcode

func minimumTimeRequiredDP(jobs []int, k int) int {
	n := len(jobs)
	m := 1 << n

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, k)
	}

	sum := make([]int, m)
	for S := 0; S < m; S++ {
		sum[S] = 0
		for i := 0; i < n; i++ {
			if (S>>i)&1 > 0 {
				sum[S] += jobs[i]
			}
		}
		dp[S][0] = sum[S]
		for j := 1; j < k; j++ {
			dp[S][j] = dp[S][j-1]
			for T := S; T > 0; T = (T - 1) & S {
				dp[S][j] = min(dp[S][j], max(dp[T][j-1], sum[S-T]))
			}
		}
	}

	return dp[m-1][k-1]
}

func minimumTimeRequired(jobs []int, k int) int {
	n, mx, answer := len(jobs), 0, 0
	for _, v := range jobs {
		answer += v
		mx = max(mx, v)
	}
	if n <= k {
		return mx
	}
	workers := make([]int, k)

	var dfs func(int, int)
	dfs = func(p, curr int) {
		if p == n {
			answer = curr
			return
		}
		M := map[int]struct{}{}
		for i := 0; i < k; i++ {
			if _, seen := M[workers[i]]; seen || workers[i]+jobs[p] >= answer {
				continue
			}
			M[workers[i]] = struct{}{}
			workers[i] += jobs[p]
			dfs(p+1, max(workers[i], curr))
			workers[i] -= jobs[p]
		}
	}

	dfs(0, 0)
	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
