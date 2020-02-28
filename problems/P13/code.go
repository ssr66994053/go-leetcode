package P13

func romanToInt(s string) int {
	sum := 0
	var last byte = ' '
	arr := []byte(s)
	for i := 0; i < len(s); i++ {
		t := arr[i]
		switch t {
		case 'I':
			sum += 1
		case 'V':
			if last == 'I' {
				sum += 3
			} else {
				sum += 5
			}
		case 'X':
			if last == 'I' {
				sum += 8
			} else {
				sum += 10
			}
		case 'L':
			if last == 'X' {
				sum += 30
			} else {
				sum += 50
			}
		case 'C':
			if last == 'X' {
				sum += 80
			} else {
				sum += 100
			}
		case 'D':
			if last == 'C' {
				sum += 300
			} else {
				sum += 500
			}
		case 'M':
			if last == 'C' {
				sum += 800
			} else {
				sum += 1000
			}
		}
		last = t
	}

	return sum
}
