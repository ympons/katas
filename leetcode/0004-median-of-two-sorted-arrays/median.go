package leetcode

// O(log(min(m,n)))
func findMedianSortedArrays(A []int, B []int) float64 {
	m, n := len(A), len(B)
	// ensure m <= n
	if m > n {
		A, B, m, n = B, A, n, m
	}
	// Given:
	// A[0], A[1], ..., A[i-1]  |  A[i], A[i+1], ..., A[m-1]
	// B[0], B[1], ..., B[j-1]  |  B[j], B[j+1], ..., B[n-1]
	//
	// Find i in [0...m], such that:
	// B[j - 1] <= A[i] and A[i - 1] <= B[j], where j = (m + n + 1) / 2 - i
	imin, imax, half, i, j := 0, m, (m+n+1)/2, 0, 0
	for imin <= imax {
		i = (imin + imax) / 2
		j = half - i
		if i < m && B[j-1] > A[i] {
			// i is too small
			imin = i + 1
		} else if i > 0 && A[i-1] > B[j] {
			// i is too big
			imax = i - 1
		} else {
			// i is perfect
			maxLeft := 0
			if i == 0 {
				maxLeft = B[j-1]
			} else if j == 0 {
				maxLeft = A[i-1]
			} else {
				maxLeft = max(A[i-1], B[j-1])
			}
			if (n+m)&1 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = B[j]
			} else if j == n {
				minRight = A[i]
			} else {
				minRight = min(A[i], B[j])
			}
			return float64(maxLeft+minRight) / 2
		}
	}
	return 0.0
}

// O((n+m)/2)
func findMedianSortedArraysOld(A []int, B []int) float64 {
	m, n := len(A), len(B)
	var mid, prevMid int
	if m > 0 {
		mid = A[0]
	} else if n > 0 {
		mid = B[0]
	}
	i, j := 0, 0
	for k := (m + n) / 2; k >= 0; k-- {
		prevMid = mid
		if j == n || (i < m && A[i] <= B[j]) {
			mid = A[i]
			i++
		} else if i == m || (j < n && B[j] <= A[i]) {
			mid = B[j]
			j++
		}
	}
	if (m+n)&1 == 1 {
		return float64(mid)
	}
	return float64(prevMid+mid) / 2
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
