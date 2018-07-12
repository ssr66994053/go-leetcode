package P110

import (
	"github.com/ssr66994053/go-leetcode/model"
	"testing"
)

func Test1(t *testing.T) {
	exp := true
	res := isBalanced(model.BuildTree([]int{3, 9, 20, 999, 999, 15, 7}))

	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}
}

func Test2(t *testing.T) {
	exp := false
	res := isBalanced(model.BuildTree([]int{1, 2, 2, 3, 3, 999, 999, 4, 4}))

	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}
}

func Test3(t *testing.T) {
	exp := false
	res := isBalanced(model.BuildTree([]int{1, 2, 2, 3, 999, 999, 3, 4, 999, 999, 999, 999, 999, 999, 4}))

	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}
}
