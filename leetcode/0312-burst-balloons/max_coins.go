package leetcode

func maxCoins(nums []int) int {
	A := make([]int, 0, len(nums)+2)
	A = append(A, 1)
	for _, x := range nums {
		if x > 0 {
			A = append(A, x)
		}
	}
	A = append(A, 1)
	n := len(A)

	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
	}
	for k := 2; k < n; k++ {
		for left := 0; left < n-k; left++ {
			right := left + k
			for i := left + 1; i < right; i++ {
				sum := A[left]*A[i]*A[right] + C[left][i] + C[i][right]
				if sum > C[left][right] {
					C[left][right] = sum
				}
			}
		}
	}
	return C[0][n-1]
}

func maxCoinsRecursive(nums []int) int {
	A := make([]int, 0, len(nums)+2)
	A = append(A, 1)
	for _, x := range nums {
		if x > 0 {
			A = append(A, x)
		}
	}
	A = append(A, 1)
	n := len(A)

	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
	}

	return dfs(A, 0, n-1, &C)
}

func dfs(nums []int, left, right int, C *[][]int) int {
	if left >= right {
		return 0
	}
	if (*C)[left][right] > 0 {
		return (*C)[left][right]
	}
	answer := 0
	for i := left + 1; i < right; i++ {
		s := nums[left]*nums[i]*nums[right] + dfs(nums, left, i, C) + dfs(nums, i, right, C)
		if answer < s {
			answer = s
		}
	}
	(*C)[left][right] = answer
	return answer
}
