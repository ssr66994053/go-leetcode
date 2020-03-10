package P88

import (
	"testing"
)

func Test1(t *testing.T) {
	exp := []int{1, 2, 2, 3, 5, 6}
	res := []int{1, 2, 3, 0, 0, 0}
	n2 := []int{2, 5, 6}
	merge(res, 3, n2, 3)

	if !equal(exp, res) {
		t.Errorf("期待 %v 实际 %v\n", exp, res)
	}
}

func Test2(t *testing.T) {
	exp := []int{1, 2, 3, 4, 5, 6}
	res := []int{4, 5, 6, 0, 0, 0}
	n2 := []int{1, 2, 3}
	merge(res, 3, n2, 3)

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
