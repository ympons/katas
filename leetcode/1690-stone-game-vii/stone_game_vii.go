package leetcode

func stoneGameVII(nums []int) int {
	n := len(nums)

	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			C[i][j] = sum[j] - sum[i] - C[i][j-1]
			if s := sum[j+1] - sum[i+1] - C[i+1][j]; s > C[i][j] {
				C[i][j] = s
			}
		}
	}

	return C[0][n-1]
}

func stoneGameVIIRecursive(nums []int) int {
	n, sum := len(nums), 0
	for _, x := range nums {
		sum += x
	}

	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
	}

	return dfs(nums, 0, n-1, sum, &C)
}

func dfs(nums []int, i, j, sum int, C *[][]int) int {
	if i >= j {
		return 0
	}
	if (*C)[i][j] > 0 {
		return (*C)[i][j]
	}
	l := sum - nums[i] - dfs(nums, i+1, j, sum-nums[i], C)
	r := sum - nums[j] - dfs(nums, i, j-1, sum-nums[j], C)
	if l > r {
		(*C)[i][j] = l
		return l
	}
	(*C)[i][j] = r
	return r
}
