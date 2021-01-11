package leetcode

func canPartitionKSubsets(nums []int, k int) bool {
	n, s, mx := len(nums), sum(nums...), max(nums...)
	bucketSum := s / k
	if n < k || s%k > 0 || mx > bucketSum {
		return false
	}
	M := map[int]struct{}{}

	var dfs func(int, int, int) bool
	dfs = func(k, start, currentSum int) bool {
		// a single bucket is always valid
		if k == 1 {
			return true
		}
		// found a valid bucket, need to check the others k-1
		if currentSum == bucketSum {
			return dfs(k-1, 0, 0)
		}
		for i := start; i < n; i++ {
			// skip if it is already visited or out of bound
			if _, visited := M[i]; visited || currentSum+nums[i] > bucketSum {
				continue
			}
			M[i] = struct{}{}
			// continue building the current bucket
			if dfs(k, i+1, currentSum+nums[i]) {
				return true
			}
			delete(M, i)
		}
		return false
	}

	return dfs(k, 0, 0)
}

func sum(nums ...int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	return s
}

func max(nums ...int) int {
	mx := 0
	for _, v := range nums {
		if mx < v {
			mx = v
		}
	}
	return mx
}
