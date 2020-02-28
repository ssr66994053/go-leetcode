package P13

import "testing"

func Test1(t *testing.T) {
	exp := 3
	res := romanToInt("III")
	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}

	exp = 4
	res = romanToInt("IV")
	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}

	exp = 9
	res = romanToInt("IX")
	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}

	exp = 58
	res = romanToInt("LVIII")
	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}

	exp = 1994
	res = romanToInt("MCMXCIV")
	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}
