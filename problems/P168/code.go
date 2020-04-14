package P168

var arr = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func convertToTitle(n int) string {
	s := ""
	if n <= 0 {
		return s
	}

	for n > 0 {
		d := n % 26

		if d == 0 {
			d = 26
			n = n - 1
		}

		s = arr[d-1] + s
		n = n / 26
	}

	return s
}
