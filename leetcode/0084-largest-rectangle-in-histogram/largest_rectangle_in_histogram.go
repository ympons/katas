package leetcode

func largestRectangleArea(heights []int) int {
	stack, n := []int{}, len(heights)

	maxArea, h := 0, 0
	for i := 0; i <= n; i++ {
		if i == n {
			h = 0
		} else {
			h = heights[i]
		}

		if len(stack) == 0 || h >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			height := heights[stack[len(stack)-1]]
			stack = stack[0 : len(stack)-1]

			width := i
			if len(stack) > 0 {
				width = i - 1 - stack[len(stack)-1]
			}

			if area := height * width; maxArea < area {
				maxArea = area
			}

			i--
		}
	}
	return maxArea
}
