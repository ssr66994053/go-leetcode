package P67

import (
	"testing"
)

func Test1(t *testing.T) {
	exp := "100"
	res := addBinary("11", "1")

	if exp != res {
		t.Errorf("期待 %s 实际 %s\n", exp, res)
	}
}
func Test2(t *testing.T) {
	exp := "10101"
	res := addBinary("1010", "1011")

	if exp != res {
		t.Errorf("期待 %s 实际 %s\n", exp, res)
	}
}

func Test3(t *testing.T) {
	exp := "1100"
	res := add([]byte("1011"), 3)

	if exp != string(res) {
		t.Errorf("期待 %s 实际 %s\n", exp, res)
	}
}

func Test4(t *testing.T) {
	exp := "110110"
	res := addBinary("100", "110010")

	if exp != res {
		t.Errorf("期待 %s 实际 %s\n", exp, res)
	}
}
