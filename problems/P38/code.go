package P38

func countAndSay(n int) string {
	t := []byte{'1'}
	for i := 2; i <= n; i++ {
		s := t[0]
		n := byte('1')
		a := []byte{}
		for j := 1; j < len(t); j++ {
			if s != t[j] {
				a = append(a, n, s)
				s = t[j]
				n = byte('1')
			} else {
				n++
			}
		}
		t = append(a, n, s)
	}

	return string(t)
}
