package P100

import (
	"github.com/ssr66994053/go-leetcode/model"
	"testing"
)

func buildTree(vals []int) *model.TreeNode {
	if vals == nil || len(vals) == 0 {
		return nil
	}
	//这里只是1层树
	root := &model.TreeNode{}
	for i, v := range vals {
		if v == 0 {
			continue
		}
		n := i % 3
		switch n {
		case 0:
			root.Val = v
			break
		case 1:
			root.Left = &model.TreeNode{Val: v}
			break
		case 2:
			root.Right = &model.TreeNode{Val: v}
		}
	}

	return root
}

func Test1(t *testing.T) {
	exp := true
	res := isSameTree(buildTree([]int{1, 2, 3}), buildTree([]int{1, 2, 3}))

	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}
}

func Test2(t *testing.T) {
	exp := false
	res := isSameTree(buildTree([]int{1, 2, 3}), buildTree([]int{1, 0, 3}))

	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}
}

func Test3(t *testing.T) {
	exp := false
	res := isSameTree(buildTree([]int{1, 2, 1}), buildTree([]int{1, 1, 2}))

	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}
}
