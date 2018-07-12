package P226

import (
	"testing"

	"github.com/ssr66994053/go-leetcode/model"
)

func Test1(t *testing.T) {
	exp := model.BuildTree([]int{4, 7, 2, 9, 6, 3, 1})
	res := invertTree(model.BuildTree([]int{4, 2, 7, 1, 3, 6, 9}))
	flag := model.CheckTree(exp, res)

	if !flag {
		t.Errorf("期待 %v 实际 %v\n", exp.StringArray(), res.StringArray())
	}
}
