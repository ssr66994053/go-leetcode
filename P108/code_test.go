package P108

import (
	"github.com/ssr66994053/go-leetcode/model"
	"testing"
)

func Test1(t *testing.T) {
	exp := model.BuildTree([]int{0, -3, 9, -10, 999, 5})
	res := sortedArrayToBST([]int{-10, -3, 0, 5, 9})

	flag := model.CheckTree(exp, res)

	if !flag {
		t.Errorf("期待 %v 实际 %v\n", exp, res)
	}
}
