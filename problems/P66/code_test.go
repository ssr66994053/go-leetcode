package P66

import (
	"testing"
)

func Test1(t *testing.T) {
	exp := []int{1, 2, 4}
	res := plusOne([]int{1, 2, 3})

	if !equal(exp, res) {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test2(t *testing.T) {
	exp := []int{4, 3, 2, 2}
	res := plusOne([]int{4, 3, 2, 1})

	if !equal(exp, res) {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test3(t *testing.T) {
	exp := []int{1, 0}
	res := plusOne([]int{9})

	if !equal(exp, res) {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func equal(s, t []int) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			return false
		}
	}

	return true
}
