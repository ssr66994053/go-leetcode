package P169

import (
	"testing"
)

func Test1(t *testing.T) {
	exp := 2
	res := majorityElement([]int{2, 1, 2})
	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}
