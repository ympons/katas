package leetcode

import "sort"

type stone struct {
	val, idx int
}

func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
	stones := make([]*stone, n)
	for i := 0; i < n; i++ {
		stones[i] = &stone{aliceValues[i] + bobValues[i], i}
	}

	sort.Slice(stones, func(i, j int) bool {
		return stones[i].val >= stones[j].val
	})

	alice, bob := 0, 0
	for i := 0; i < n; i++ {
		idx := stones[i].idx
		if i%2 == 0 {
			alice += aliceValues[idx]
		} else {
			bob += bobValues[idx]
		}
	}
	if alice > bob {
		return 1
	}
	if alice < bob {
		return -1
	}
	return 0
}
