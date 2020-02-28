package P14

import "testing"

func Test1(t *testing.T) {
	exp := "fl"
	res := longestCommonPrefix([]string{"flower", "flow", "flight"})
	if exp != res {
		t.Errorf("期待 %s 实际 %s\n", exp, res)
	}

	exp = ""
	res = longestCommonPrefix([]string{"dog", "racecar", "car"})
	if exp != res {
		t.Errorf("期待 %s 实际 %s\n", exp, res)
	}

	exp = ""
	res = longestCommonPrefix([]string{})
	if exp != res {
		t.Errorf("期待 %s 实际 %s\n", exp, res)
	}

	exp = "a"
	res = longestCommonPrefix([]string{"a"})
	if exp != res {
		t.Errorf("期待 %s 实际 %s\n", exp, res)
	}

	exp = "c"
	res = longestCommonPrefix([]string{"c", "c"})
	if exp != res {
		t.Errorf("期待 %s 实际 %s\n", exp, res)
	}

}
