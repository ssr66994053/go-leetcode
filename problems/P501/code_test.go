package P501

import (
	"testing"

	"github.com/ssr66994053/go-leetcode/model"
	"sort"
)

func Test1(t *testing.T) {
	exp := []int{2}
	res := findMode(model.BuildTree([]int{1, 999, 2, 999, 999, 2, 999}))

	if len(exp) != len(res) {
		t.Errorf("期待 %v 实际 %v\n", exp, res)
	}

	sort.Ints(exp)
	sort.Ints(res)
	for i, v := range exp {
		if v != res[i] {
			t.Errorf("期待 %v 实际 %v\n", exp, res)
		}
	}
}
