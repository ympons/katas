package leetcode

func countGoodRectangles(rectangles [][]int) int {
	m, k, mx := map[int]int{}, 0, 0
	for _, r := range rectangles {
		k = min(r[0], r[1])
		m[k]++
		if mx < k {
			mx = k
		}
	}
	return m[mx]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
