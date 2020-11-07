package leetcode

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		ni := n1 + n2
		n1, n2 = n2, ni
	}
	return n2
}

func climbStairsDP(n int) int {
	fib := make([]int, n+1)
	fib[0], fib[1] = 1, 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}
