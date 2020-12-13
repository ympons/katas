package leetcode

import (
	"math/bits"
	"sort"
)

const maxInt = int(1e9)

// Using Submask Enumeration:
// https://cp-algorithms.com/algebra/all-submasks.html
// Time: O(n*2^n + 3^n)
func minimumIncompatibility(nums []int, k int) int {
	n := len(nums)
	if n < k {
		return -1
	}
	m := n / k
	diff := make([]int, 1<<n)
	cur := make([]int, 0, n)
	// O(n*2^n)
	for S := 1; S < (1 << n); S++ {
		cur = cur[0:0]
		for i := 0; i < n; i++ {
			if (S>>i)&1 > 0 {
				cur = append(cur, nums[i])
			}
		}
		sort.Ints(cur)
		repeated := false
		for i := 1; i < len(cur); i++ {
			if cur[i] == cur[i-1] {
				repeated = true
				break
			}
		}
		if repeated {
			diff[S] = maxInt
		} else {
			diff[S] = cur[len(cur)-1] - cur[0]
		}
	}
	// Submask Enumeration ~ O(3^n)
	dp := make([]int, 1<<n)
	for S := 1; S < (1 << n); S++ {
		dp[S] = maxInt
		for T := S; T > 0; T = (T - 1) & S {
			if bits.OnesCount(uint(T)) == m {
				if sum := diff[T] + dp[S^T]; sum < dp[S] {
					dp[S] = sum
				}
			}
		}
	}
	answer := dp[(1<<n)-1]
	if answer == maxInt {
		return -1
	}
	return answer
}
