package P28

func strStr(haystack string, needle string) int {
	// return strings.Index(haystack, needle)

	if haystack == needle {
		return 0
	}

	if len(haystack) < len(needle) {
		return -1
	}

	s := []byte(haystack)
	t := []byte(needle)

	idx := -1
	for i := 0; i < len(s); i++ {
		f := false
		for j := 0; j < len(t) && i+j < len(s); j++ {
			if s[i+j] != t[j] {
				f = true
				break
			}
		}

		if f {
			continue
		}

		if i+len(t) <= len(s) {
			idx = i
		}

		break
	}

	return idx
}
