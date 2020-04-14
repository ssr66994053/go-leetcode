package P167

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for i <= j {
		a := numbers[i]
		b := numbers[j]
		if target < a+b {
			j--
			continue
		}
		if target > a+b {
			i++
			continue
		}

		return []int{i + 1, j + 1}
	}

	return nil
}
