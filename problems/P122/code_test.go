package P121

import (
	"testing"
)

func Test1(t *testing.T) {
	exp := 7
	res := maxProfit([]int{7, 1, 5, 3, 6, 4})

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test2(t *testing.T) {
	exp := 4
	res := maxProfit([]int{1, 2, 3, 4, 5})

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test3(t *testing.T) {
	exp := 2
	res := maxProfit([]int{2, 4, 1})

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}
