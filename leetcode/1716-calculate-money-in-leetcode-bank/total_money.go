package leetcode

// Time: O(n), Space: O(1)
func totalMoney(n int) int {
	sum, week, count := 1, 1, 2
	for i := 1; i < n; i++ {
		if i%7 == 0 {
			week++
			count = week
		}
		sum += count
		count++
	}
	return sum
}

// Time: O(n), Space: O(n)
func totalMoneyDP(n int) int {
	total, sum := make([]int, n), 1
	total[0] = 1
	for i := 1; i < n; i++ {
		if i%7 != 0 {
			total[i] = total[i-1] + 1
		} else {
			total[i] = total[i-7] + 1
		}
		sum += total[i]
	}
	return sum
}
