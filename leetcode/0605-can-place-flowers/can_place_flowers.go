package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	m, placed := len(flowerbed), 0
	for i := 0; i < m; i++ {
		if flowerbed[i] == 1 {
			continue
		}
		if (i == 0 || flowerbed[i-1] == 0) && (i == m-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			placed++
		}
	}
	return placed >= n
}
