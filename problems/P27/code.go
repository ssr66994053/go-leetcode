package P27

func removeElement(nums []int, val int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	if l == 1 {
		if nums[0] == val {
			return 0
		}
		return 1
	}

	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}

		i++
	}

	return len(nums)
}
