package P125

import (
	"testing"
)

func Test1(t *testing.T) {
	exp := true
	res := isPalindrome("A man, a plan, a canal: Panama")
	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}
}

func Test2(t *testing.T) {
	exp := true
	res := isPalindrome("aa")
	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}
}
