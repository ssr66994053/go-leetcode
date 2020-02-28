package P9

import "testing"

func Test1(t *testing.T) {
	res := isPalindrome(121)
	if !res {
		t.Error("期待 true 实际 false\n")
	}

}
