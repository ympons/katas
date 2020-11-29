package leetcode

func canReachDFS(nums []int, start int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	return dfs(append([]int{}, nums...), start, n)
}

func dfs(nums []int, i, n int) bool {
	if 0 <= i && i < n && nums[i] >= 0 {
		if nums[i] == 0 {
			return true
		}
		nums[i] = -nums[i]
		return dfs(nums, i-nums[i], n) || dfs(nums, i+nums[i], n)
	}
	return false
}

var visited map[int]struct{}

func canReach(nums []int, start int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	visited = map[int]struct{}{}
	return bfs(nums, start, n)
}

func bfs(nums []int, start, n int) bool {
	q, i := []int{start}, 0
	for len(q) > 0 {
		i = q[0]
		if nums[i] == 0 {
			return true
		}
		q = q[1:len(q)]
		if _, ok := visited[i-nums[i]]; !ok && i-nums[i] >= 0 {
			visited[i-nums[i]] = struct{}{}
			q = append(q, i-nums[i])
		}
		if _, ok := visited[i+nums[i]]; !ok && i+nums[i] < n {
			visited[i+nums[i]] = struct{}{}
			q = append(q, i+nums[i])
		}
	}
	return false
}
