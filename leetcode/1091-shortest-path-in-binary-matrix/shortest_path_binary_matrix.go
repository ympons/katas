package leetcode

// Point represents a grid position
type Point struct {
	x, y int
}

// Time: O(N^2), Space: O(N^2)
func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	PS := []Point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	M, queue, steps := map[Point]struct{}{}, []Point{{0, 0}}, 0
	M[queue[0]] = struct{}{}
	for len(queue) > 0 {
		steps++
		for level := len(queue); level > 0; level-- {
			point := queue[0]
			queue = queue[1:]
			if point.x == n-1 && point.y == n-1 {
				return steps
			}
			for _, p := range PS {
				newPoint := Point{point.x + p.x, point.y + p.y}
				_, visited := M[newPoint]
				if visited || newPoint.x < 0 || newPoint.x >= n || newPoint.y < 0 || newPoint.y >= n || grid[newPoint.x][newPoint.y] == 1 {
					continue
				}
				M[newPoint] = struct{}{}
				queue = append(queue, newPoint)
			}
		}
	}

	return -1
}
