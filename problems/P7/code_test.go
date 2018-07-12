package P7

import "testing"

func Test1(t *testing.T) {
	exp := 321
	res := reverse(123)
	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test2(t *testing.T) {
	exp := -321
	res := reverse(-123)
	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test3(t *testing.T) {
	exp := 21
	res := reverse(120)
	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test4(t *testing.T) {
	exp := 0
	res := reverse(1534236469)
	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}
