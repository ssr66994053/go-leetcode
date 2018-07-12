package P107

import (
	"github.com/ssr66994053/go-leetcode/model"
	"testing"
)

func Test1(t *testing.T) {
	exp := [][]int{{15, 7}, {9, 20}, {3}}
	res := levelOrderBottom(model.BuildTree([]int{3, 9, 20, 999, 999, 15, 7}))

	flag := false
	for i, v := range exp {
		rV := res[i]
		iFlag := false
		for j, val := range v {
			if val != rV[j] {
				iFlag = true
				break
			}
		}
		if iFlag {
			flag = true
			break
		}
	}
	if flag {
		t.Errorf("期待 %v 实际 %v\n", exp, res)
	}
}
