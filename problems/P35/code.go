package P35

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target > nums[i] {
			continue
		}

		return i
	}

	return len(nums)
}
