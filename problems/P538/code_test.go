package P538

import (
	"testing"

	"github.com/ssr66994053/go-leetcode/model"
)

func Test1(t *testing.T) {
	exp := model.BuildTree([]int{18, 20, 13})
	res := convertBST(model.BuildTree([]int{5, 2, 13}))

	if !model.CheckTree(exp, res) {
		t.Errorf("期待 %v 实际 %v\n", exp, res)
	}
}
