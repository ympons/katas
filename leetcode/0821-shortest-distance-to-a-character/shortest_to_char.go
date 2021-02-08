package leetcode

// Time: O(N), Space: O(N)
func shortestToChar(s string, c byte) []int {
	n := len(s)
	answer := make([]int, n)
	prev := -2 * n
	for i := range s {
		if s[i] == c {
			prev = i
		}
		answer[i] = i - prev
	}
	prev = 2 * n
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			prev = i
		}
		answer[i] = min(answer[i], prev-i)
	}
	return answer
}

// Time: O(N), Space: O(2*N)
func shortestToChar2(s string, c byte) []int {
	n := len(s)
	D, q := make([]int, n), []int{}
	for i := range s {
		if s[i] == c {
			q = append(q, i)
			D[i] = 0
		} else {
			D[i] = n
		}
	}
	j := 0
	for i := range s {
		if s[i] == c {
			j++
		} else if j < len(q) {
			D[i] = min(D[i], abs(i-q[j]))
		}
		if j-1 >= 0 {
			D[i] = min(D[i], abs(i-q[j-1]))
		}
	}
	return D
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
