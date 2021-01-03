package leetcode

const module = int(1e9 + 7)

func waysToSplit(nums []int) int {
	n := len(nums)

	S := make([]int, n)
	S[0] = nums[0]
	for i := 1; i < n; i++ {
		S[i] = S[i-1] + nums[i]
	}

	answer := 0
	for i := 1; i < n-1; i++ {
		if l, r := search(S, i, false), search(S, i, true); l != -1 && r != -1 {
			answer = (answer + (r - l + 1)) % module
		}
	}

	return answer
}

func search(S []int, index int, upper bool) int {
	n, pos := len(S), -1
	l, r := index, n-2
	// SL <= SM <= SR
	SL, SM, SR := S[index-1], 0, 0
	for l <= r {
		m := l + (r-l)/2
		SM, SR = S[m]-SL, S[n-1]-S[m]
		if SL <= SM && SM <= SR {
			pos = m
			if upper {
				l = m + 1
			} else {
				r = m - 1
			}
		} else if SL > SM {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return pos
}
