package P167

import (
	"testing"
)

func Test1(t *testing.T) {
	exp := []int{1, 2}
	res := twoSum([]int{2, 7, 11, 15}, 9)
	if !equal(exp, res) {
		t.Errorf("期待 %v 实际 %v\n", exp, res)
	}
}

func equal(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
