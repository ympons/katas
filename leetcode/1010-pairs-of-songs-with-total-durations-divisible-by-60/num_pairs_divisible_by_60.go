package leetcode

func numPairsDivisibleBy60(time []int) int {
	m, pairs := map[int]int{}, 0
	for _, x := range time {
		xmod := x % 60
		if p, ok := m[(60-xmod)%60]; ok {
			pairs += p
		}
		m[xmod]++
	}
	return pairs
}
