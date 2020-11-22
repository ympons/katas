package leetcode

func minCostToMoveChips(position []int) int {
	even, odd := 0, 0
	for _, p := range position {
		if p&1 == 1 {
			odd++
		} else {
			even++
		}
	}
	if even < odd {
		return even
	}
	return odd
}
