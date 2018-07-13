package P404

import (
	"testing"

	"github.com/ssr66994053/go-leetcode/model"
)

func Test1(t *testing.T) {
	exp := 24
	res := sumOfLeftLeaves(model.BuildTree([]int{3, 9, 20, 999, 999, 15, 7}))

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test2(t *testing.T) {
	exp := 0
	res := sumOfLeftLeaves(model.BuildTree([]int{1}))

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}
