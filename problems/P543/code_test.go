package P543

import (
	"testing"

	"github.com/ssr66994053/go-leetcode/model"
)

func Test1(t *testing.T) {
	exp := 3
	res := diameterOfBinaryTree(model.BuildTree([]int{1, 2, 3, 4, 5, 999, 999}))

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}
