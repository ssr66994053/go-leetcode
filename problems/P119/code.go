package P119

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	if rowIndex == 1 {
		return []int{1, 1}
	}

	last := getRow(rowIndex - 1)
	a := make([]int, len(last)+1, len(last)+1)
	for j := 0; j < len(a); j++ {
		if j == 0 || j == len(a)-1 {
			a[j] = 1
			continue
		}

		a[j] = last[j-1] + last[j]
	}

	return a
}
