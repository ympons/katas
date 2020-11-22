package leetcode

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	neighbors := make([]map[int]struct{}, n)
	for i := 0; i < n; i++ {
		neighbors[i] = map[int]struct{}{}
	}

	for _, edge := range edges {
		neighbors[edge[0]][edge[1]] = struct{}{}
		neighbors[edge[1]][edge[0]] = struct{}{}
	}

	leaves := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if len(neighbors[i]) == 1 {
			leaves = append(leaves, i)
		}
	}

	for n > 2 {
		n -= len(leaves)
		newLeaves := make([]int, 0, n)
		for _, leaf := range leaves {
			for i := range neighbors[leaf] {
				delete(neighbors[i], leaf)
				if len(neighbors[i]) == 1 {
					newLeaves = append(newLeaves, i)
				}
			}
		}
		leaves = newLeaves
	}

	return leaves
}
