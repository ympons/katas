package leetcode

func constructDistancedSequence(n int) []int {
	if n == 1 {
		return []int{1}
	}

	size := 2*n - 1
	answer := make([]int, size)

	var dfs func(int, int) bool
	dfs = func(mask int, p int) bool {
		// if all the values have been assigned we found the answer
		if p == size {
			return true
		}
		// if the value is already assigned move to the next index
		if answer[p] != 0 {
			return dfs(mask, p+1)
		}
		// iterate the numbers in descending order to ensure a lexicographically largest sequence
		for i := n; i >= 1; i-- {
			// skip i if is already in used or its second occurrence is out of bound or cannot be filled
			if mask&(1<<i) != 0 || (i > 1 && (p+i >= size || answer[p+i] != 0)) {
				continue
			}
			// fill i and its second occurrence (if it is possible)
			answer[p] = i
			if i > 1 {
				answer[p+i] = i
			}
			// mark i in the mask and move to the next index; stop if the sequence has been found
			if dfs(mask|(1<<i), p+1) {
				return true
			}
			// backtrack the filling
			answer[p] = 0
			if i > 1 {
				answer[p+i] = 0
			}
		}
		return false
	}

	dfs(0, 0)
	return answer
}
