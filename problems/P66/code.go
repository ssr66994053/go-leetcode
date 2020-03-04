package P66

func plusOne(digits []int) []int {
	l := len(digits)
	t := digits[l-1]
	t++
	digits[l-1] = t

	if t == 10 {
		digits[l-1] = 0
		if l == 1 {
			return []int{1, 0}
		}

		return append(plusOne(digits[:l-1]), 0)
	} else {
		return digits
	}
}
