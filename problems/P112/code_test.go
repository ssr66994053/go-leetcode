package P112

import (
	"github.com/ssr66994053/go-leetcode/model"
	"testing"
)

func Test1(t *testing.T) {
	exp := true
	res := hasPathSum(model.BuildTree([]int{5, 4, 8, 11, 999, 13, 4, 7, 2, 999, 999, 999, 1}), 22)

	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}
}

func Test2(t *testing.T) {
	exp := true
	res := hasPathSum(model.BuildTree([]int{-2, 999, -3}), -5)

	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}
}
