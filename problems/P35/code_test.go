package P35

import (
	"testing"
)

func Test1(t *testing.T) {
	exp := 2
	res := searchInsert([]int{1, 3, 5, 6}, 5)

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}
