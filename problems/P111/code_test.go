package P111

import (
	"github.com/ssr66994053/go-leetcode/model"
	"testing"
)

func Test1(t *testing.T) {
	exp := 2
	res := minDepth(model.BuildTree([]int{3, 9, 20, 999, 999, 15, 7}))

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test2(t *testing.T) {
	exp := 2
	res := minDepth(model.BuildTree([]int{1, 2, 999}))

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test3(t *testing.T) {
	exp := 3
	res := minDepth(model.BuildTree([]int{1, 2, 3, 4, 999, 999, 5}))

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test4(t *testing.T) {
	exp := 1
	res := minDepth(model.BuildTree([]int{0}))

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test5(t *testing.T) {
	exp := 5
	res := minDepth(model.BuildTree([]int{1, 2, 999, 3, 999, 999, 999, 4, 999, 999, 999, 999, 999, 999, 999,
		5, 999, 999, 999, 999, 999, 999, 999, 999, 999, 999, 999, 999, 999, 999, 999}))

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}
