package P168

import (
	"testing"
)

func Test1(t *testing.T) {
	exp := "A"
	res := convertToTitle(1)
	if exp != res {
		t.Errorf("期待 %s 实际 %s\n", exp, res)
	}
}

func Test2(t *testing.T) {
	exp := "Z"
	res := convertToTitle(26)
	if exp != res {
		t.Errorf("期待 %s 实际 %s\n", exp, res)
	}
}

func Test3(t *testing.T) {
	exp := "AA"
	res := convertToTitle(27)
	if exp != res {
		t.Errorf("期待 %s 实际 %s\n", exp, res)
	}
}

func Test4(t *testing.T) {
	exp := "ZZ"
	res := convertToTitle(52)
	if exp != res {
		t.Errorf("期待 %s 实际 %s\n", exp, res)
	}
}
