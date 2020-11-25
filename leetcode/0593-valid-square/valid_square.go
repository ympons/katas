package leetcode

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	points := [][]int{p1, p2, p3, p4}
	return helper(points, 0)
}

func helper(points [][]int, p int) bool {
	if p == 4 {
		return valid(points)
	}
	check := false
	for i := p; i < 4; i++ {
		points[i], points[p] = points[p], points[i]
		check = check || helper(points, p+1)
		points[i], points[p] = points[p], points[i]
	}
	return check
}

func distance(p1, p2 []int) int {
	deltaX := p1[0] - p2[0]
	deltaY := p1[1] - p2[1]
	return deltaX*deltaX + deltaY*deltaY
}

func valid(points [][]int) bool {
	d12 := distance(points[0], points[1])
	d14 := distance(points[0], points[3])
	d32 := distance(points[2], points[1])
	d34 := distance(points[2], points[3])
	d13 := distance(points[0], points[2])
	d24 := distance(points[1], points[3])
	return d12 > 0 && d12 == d14 && d32 == d34 && d12 == d32 && d13 == d12+d32 && d24 == d12+d14
}
