package leetcode

func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{buildings[0][0], buildings[0][2]}, {buildings[0][1], 0}}
	}
	middle := n / 2
	left := getSkyline(buildings[0:middle])
	right := getSkyline(buildings[middle:n])
	return merge(left, right)
}

func merge(left, right [][]int) [][]int {
	output := [][]int{}
	pos, h1, h2 := 0, 0, 0
	i, j, n, m := 0, 0, len(left), len(right)
	for i < n && j < m {
		if left[i][0] < right[j][0] {
			pos, h1 = left[i][0], left[i][1]
			i++
		} else if left[i][0] > right[j][0] {
			pos, h2 = right[j][0], right[j][1]
			j++
		} else {
			h1, h2 = left[i][1], right[j][1]
			pos = right[j][0]
			i++
			j++
		}
		maxH := max(h1, h2)
		if len(output) == 0 || output[len(output)-1][1] != maxH {
			output = append(output, []int{pos, maxH})
		}
	}
	if i < n {
		output = append(output, left[i:n]...)
	}
	if j < m {
		output = append(output, right[j:m]...)
	}
	return output
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
