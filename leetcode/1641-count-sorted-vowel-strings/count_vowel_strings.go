package leetcode

// Time: O(N), Space: O(N)
func countVowelStringsDP(n int) int {
	dp := []int{0, 1, 1, 1, 1, 1}
	for ; n > 0; n-- {
		for i := 1; i <= 5; i++ {
			dp[i] += dp[i-1]
		}
	}
	return dp[5]
}

// Combinations with repetitions: (n + k - 1)! / k! * (n - 1)!
// Time: O(1), Space: O(1)
func countVowelStrings(n int) int {
	return (n + 1) * (n + 2) * (n + 3) * (n + 4) / 24
}
