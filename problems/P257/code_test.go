package P257

import (
	"testing"

	"github.com/ssr66994053/go-leetcode/model"
)

func Test1(t *testing.T) {
	exp := []string{"1->2->5", "1->3"}
	res := binaryTreePaths(model.BuildTree([]int{1, 2, 3, 999, 5, 999, 999}))

	for i, p := range exp {
		if p != res[i] {
			t.Errorf("期待 %s 实际 %s\n", p, res[i])
		}
	}
}
