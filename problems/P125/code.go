package P125

func isPalindrome(s string) bool {
	arr := []byte(s)
	l := len(arr)
	if l == 0 {
		return true
	}

	i := 0
	j := l - 1
	for i <= j {
		x := toLower(arr[i])
		if !isAN(x) {
			i++
			continue
		}

		y := toLower(arr[j])
		if !isAN(y) {
			j--
			continue
		}

		if x != y {
			return false
		}

		i++
		j--
	}

	return true
}

func toLower(b byte) byte {
	if b >= byte('A') && b <= byte('Z') {
		b += 32
	}

	return b
}

func isAN(b byte) bool {
	if (b >= byte('a') && b <= byte('z')) || (b >= byte('0') && b <= byte('9')) {
		return true
	}

	return false
}
