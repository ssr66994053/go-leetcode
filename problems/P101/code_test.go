package P101

import (
	"github.com/ssr66994053/go-leetcode/model"
	"testing"
)

func Test1(t *testing.T) {
	exp := true
	res := isSymmetric(model.BuildTree([]int{1, 2, 2, 3, 4, 4, 3}))

	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}
}

func Test2(t *testing.T) {
	exp := false
	res := isSymmetric(model.BuildTree([]int{1, 2, 2, 0, 3, 0, 3}))

	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}
}

func Test3(t *testing.T) {
	exp := true
	res := isSymmetric(model.BuildTree([]int{1, 2, 2, 0, 3, 3, 0}))

	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}
}
