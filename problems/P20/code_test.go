package P20

import "testing"

func Test1(t *testing.T) {
	exp := true
	res := isValid("()")
	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}

	exp = true
	res = isValid("()[]{}")
	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}

	exp = false
	res = isValid("(]")
	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}

	exp = false
	res = isValid("([)]")
	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}

	exp = true
	res = isValid("{[]}")
	if exp != res {
		t.Errorf("期待 %t 实际 %t\n", exp, res)
	}

}
