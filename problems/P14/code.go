package P14

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	i := 0
	for i < len(strs[0]) {
		end := false
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				end = true
				break
			}

			if strs[0][i] != strs[j][i] {
				end = true
				break
			}
		}

		if end {
			return strs[0][0:i]
		}

		i++
	}

	return strs[0][0:i]
}
