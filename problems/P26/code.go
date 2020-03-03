package P26

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}

	tmp := nums[0]
	i := 1
	for i < len(nums) {
		t := nums[i]
		if tmp != t {
			tmp = t
			i++
			continue
		}

		nums = append(nums[:i], nums[i+1:]...)
	}

	return len(nums)
}
