package leetcode

func findBall(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	answer := make([]int, n)
	for i := 0; i < n; i++ {
		answer[i] = dfs(grid, 0, i, m, n)
	}
	return answer
}

func dfs(grid [][]int, r, c, m, n int) int {
	if r == m {
		return c
	}

	if grid[r][c] == 1 {
		if c+1 >= n || grid[r][c+1] == -1 {
			return -1
		}
		return dfs(grid, r+1, c+1, m, n)
	}

	if c-1 < 0 || grid[r][c-1] == 1 {
		return -1
	}

	return dfs(grid, r+1, c-1, m, n)
}
