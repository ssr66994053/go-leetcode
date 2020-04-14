package P169

func majorityElement(nums []int) int {
	m := 0
	n := 0

	for i := 0; i < len(nums); i++ {
		if m == 0 {
			m = nums[i]
			n = 1
			continue
		}

		if m == nums[i] {
			n++
		} else {
			n--
			if n == 0 {
				m = 0
			}
		}
	}

	return m
}
