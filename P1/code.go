package P1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for j, n := range nums {
		c := target - n
		i, ok := m[c]
		if ok {
			return []int{i, j}
		}
		m[n] = j
	}

	return []int{}
}
