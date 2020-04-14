package P136

import (
	"testing"
)

func Test1(t *testing.T) {
	exp := 1
	res := singleNumber([]int{2, 2, 1})
	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}
