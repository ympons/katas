package leetcode

var visited map[int]struct{}

func canPartitionKSubsets(nums []int, k int) bool {
	n := len(nums)
	if n < k {
		return false
	}

	sum, max := 0, 0
	for _, x := range nums {
		sum += x
		if x > max {
			max = x
		}
	}

	if sum%k > 0 {
		return false
	}

	bucketSum := sum / k
	if max > bucketSum {
		return false
	}

	visited = map[int]struct{}{}

	return helper(nums, k, 0, 0, bucketSum)
}

func helper(nums []int, k, start, currentSum, bucketSum int) bool {
	// a single bucket is always valid
	if k == 1 {
		return true
	}
	// found a valid bucket, need to check the others k-1
	if currentSum == bucketSum {
		return helper(nums, k-1, 0, 0, bucketSum)
	}
	n := len(nums)
	for i := start; i < n; i++ {
		// if it hasn't been visited try adding it to the current bucket
		newSum := currentSum + nums[i]
		if _, ok := visited[i]; !ok && newSum <= bucketSum {
			visited[i] = struct{}{}
			// continue building the current bucket
			if helper(nums, k, i+1, newSum, bucketSum) {
				return true
			}
			delete(visited, i)
		}
	}
	return false
}
