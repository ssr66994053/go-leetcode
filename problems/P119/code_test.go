package P119

import (
	"testing"
)

func Test1(t *testing.T) {
	exp := []int{1, 3, 3, 1}
	res := getRow(3)

	if !equal(exp, res) {
		t.Errorf("期待 %v 实际 %v\n", exp, res)
	}
}

func equal(a1, a2 []int) bool {
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}

	return true
}
