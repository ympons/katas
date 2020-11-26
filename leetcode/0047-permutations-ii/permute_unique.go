package leetcode

var (
	output  [][]int
	counter map[int]int
)

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	output, counter = [][]int{}, map[int]int{}
	for _, x := range nums {
		counter[x]++
	}
	helper(n, make([]int, 0, n))
	return output
}

func helper(n int, current []int) {
	if len(current) == n {
		tmp := make([]int, n)
		copy(tmp, current)
		output = append(output, tmp)
		return
	}

	for num, count := range counter {
		if count > 0 {
			current = append(current, num)
			counter[num]--
			helper(n, current)
			counter[num] = count
			current = current[0 : len(current)-1]
		}
	}
}
