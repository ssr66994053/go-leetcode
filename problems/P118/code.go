package P118

func generate(numRows int) [][]int {
	arr := make([][]int, numRows, numRows)
	for i := 0; i < numRows; i++ {
		if i == 0 {
			arr[i] = []int{1}
			continue
		}

		a := make([]int, i+1, i+1)
		t := arr[i-1]
		arr[i] = a
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				a[j] = 1
				continue
			}

			a[j] = t[j-1] + t[j]
		}
	}

	return arr
}
