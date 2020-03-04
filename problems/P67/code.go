package P67

const one = byte('1')
const zero = byte('0')

func addBinary(a string, b string) string {
	l := len(a)
	s := len(b)
	x := []byte(a)
	y := []byte(b)
	if l < s {
		s = l
		l = len(b)
		x = y
		y = []byte(a)
	}

	for i := 0; i < s; i++ {
		if x[l-i-1] == one && y[s-i-1] == one {
			x = add(x, l-i-1)
			l = len(x)
		} else if x[l-i-1] == zero && y[s-i-1] == zero {

		} else {
			x[l-i-1] = one
		}
	}

	return string(x)
}

func add(x []byte, i int) []byte {
	if i == 0 {
		return append([]byte{one, zero}, x[1:]...)
	}

	if x[i] == one {
		return append(add(x[:i], i-1), append([]byte{zero}, x[i+1:]...)...)
	}

	x[i] = one

	return x
}
