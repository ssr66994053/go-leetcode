package P104

import (
	"github.com/ssr66994053/go-leetcode/model"
	"testing"
)

func Test1(t *testing.T) {
	exp := 3
	res := maxDepth(model.BuildTree([]int{3, 9, 20, 0, 0, 15, 7}))

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}
