package leetcode

func minJumps(arr []int) int {
	n, graph := len(arr), map[int][]int{}
	for i := 0; i < n; i++ {
		graph[arr[i]] = append(graph[arr[i]], i)
	}

	queue, visited := make([]int, 0, n), map[int]struct{}{}
	queue = append(queue, 0)
	visited[0] = struct{}{}

	steps := 0
	for len(queue) > 0 {
		for lvl := len(queue); lvl > 0; lvl-- {
			i := queue[0]
			queue = queue[1:]
			if i == n-1 {
				return steps
			}
			neighbors := graph[arr[i]]
			neighbors = append(neighbors, i-1)
			neighbors = append(neighbors, i+1)
			for _, j := range neighbors {
				if _, ok := visited[j]; 0 <= j && j < n && !ok {
					queue = append(queue, j)
					visited[j] = struct{}{}
				}
			}
			graph[arr[i]] = []int{}
		}
		steps++
	}
	return -1
}
