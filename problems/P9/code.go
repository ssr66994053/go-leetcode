package P9

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	a := x
	y := 0
	for a > 0 {
		y = y*10 + a%10
		a /= 10
	}

	return x == y
}
