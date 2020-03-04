package P58

func lengthOfLastWord(s string) int {
	t := byte(' ')
	arr := []byte(s)
	n := 0
	for i := 0; i < len(arr); i++ {
		if t == arr[i] {
			n = 0
		} else {
			n++
		}
	}

	return n
}
