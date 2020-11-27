package leetcode

func canPartition(nums []int) bool {
	n := len(nums)

	if n == 1 {
		return false
	}

	sum := 0
	for _, x := range nums {
		sum += x
	}

	if sum&1 == 1 {
		return false
	}

	bucketSum := sum / 2

	D := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		D[i] = make([]bool, bucketSum+1)
	}

	D[0][0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j <= bucketSum; j++ {
			if j < nums[i-1] {
				D[i][j] = D[i-1][j]
			} else {
				D[i][j] = D[i-1][j] || D[i-1][j-nums[i-1]]
			}
		}
	}
	return D[n][bucketSum]
}
