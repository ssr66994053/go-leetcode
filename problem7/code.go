package problem7

import (
	"strconv"
	"strings"
)

func reverse(x int) int {
	if x == 0 {
		return x
	}

	s := strconv.Itoa(x)
	pre := ""
	arr := []byte(s)
	if strings.Index(s, "-") == 0 {
		pre = s[0:1]
		arr = []byte(s[1:])
	}
	b := make([]byte, 0)

	for i := len(arr) - 1; i >= 0; i-- {
		b = append(b, arr[i])
	}

	i, err := strconv.ParseInt(pre+string(b), 10, 32)
	if err != nil {
		return 0
	}

	return int(i)
}
