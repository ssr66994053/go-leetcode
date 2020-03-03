package P28

import (
	"testing"
)

func Test1(t *testing.T) {
	exp := 2
	res := strStr("hello", "ll")

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test2(t *testing.T) {
	exp := -1
	res := strStr("aaaaa", "bba")

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test3(t *testing.T) {
	exp := 0
	res := strStr("", "")

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test4(t *testing.T) {
	exp := -1
	res := strStr("mississippi", "issipi")

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}

func Test5(t *testing.T) {
	exp := 1
	res := strStr("missi", "issi")

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}
