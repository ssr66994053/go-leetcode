package P121

import (
	"testing"
)

func Test1(t *testing.T) {
	exp := 5
	res := maxProfit([]int{7, 1, 5, 3, 6, 4})

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}
