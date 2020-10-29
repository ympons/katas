package leetcode

// O(n^2)
func longestPalindrome(s string) string {
	n := len(s)
	d1, d2 := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		d1[i] = 1
		for 0 <= i-d1[i] && i+d1[i] < n && s[i-d1[i]] == s[i+d1[i]] {
			d1[i]++
		}
		d2[i] = 0
		for 0 <= i-d2[i]-1 && i+d2[i] < n && s[i-d2[i]-1] == s[i+d2[i]] {
			d2[i]++
		}
	}

	d1Idx, d1Max := 0, 0
	for k, v := range d1 {
		if v > d1Max {
			d1Max, d1Idx = v, k
		}
	}

	d2Idx, d2Max := 0, 0
	for k, v := range d2 {
		if v > d2Max {
			d2Max, d2Idx = v, k
		}
	}

	if d1Max > d2Max {
		return s[d1Idx-(d1Max-1) : d1Idx+d1Max]
	}

	return s[d2Idx-d2Max : d2Idx+d2Max]
}
