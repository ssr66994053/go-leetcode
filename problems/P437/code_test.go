package P437

import (
	"testing"

	"github.com/ssr66994053/go-leetcode/model"
)

func Test1(t *testing.T) {
	exp := 3
	res := pathSum(model.BuildTree([]int{10, 5, -3, 3, 2, 999, 11, 3, -2, 999, 1, 999, 999}), 8)

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test2(t *testing.T) {
	exp := 1
	res := pathSum(model.BuildTree([]int{-2, 999, -3}), -5)

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test3(t *testing.T) {
	exp := 3
	res := pathSum(model.BuildTree([]int{10, 5, -3, 3, 2, 999, 11, 3, -2, 999, 1, 999, 999, 999, 999}), 8)

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test4(t *testing.T) {
	exp := 4
	res := pathSum(model.BuildTree([]int{1, -2, -3, 1, 3, -2, 999, -1, 999, 999, 999, 999, 999, 999, 999}), -1)

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}
