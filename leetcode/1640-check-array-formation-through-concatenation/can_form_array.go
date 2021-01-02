package leetcode

func canFormArray(arr []int, pieces [][]int) bool {
	mark := map[int]struct{}{}
	for _, n := range arr {
		mark[n] = struct{}{}
	}
	start := map[int]int{}
	for pos, piece := range pieces {
		if len(piece) > 0 {
			start[piece[0]] = pos
			for _, n := range piece {
				if _, ok := mark[n]; !ok {
					return false
				}
				delete(mark, n)
			}
		}
	}
	if len(mark) > 0 {
		return false
	}

	i := 0
	for i < len(arr) {
		pos, ok := start[arr[i]]
		if !ok {
			return false
		}
		piece, j := pieces[pos], 0
		for j < len(piece) && i < len(arr) {
			if piece[j] != arr[i] {
				return false
			}
			i++
			j++
		}
	}

	return true
}
