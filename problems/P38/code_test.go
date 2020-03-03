package P38

import (
	"testing"
)

func Test1(t *testing.T) {
	exp := "1211"
	res := countAndSay(4)

	if exp != res {
		t.Errorf("期待 %s 实际 %s\n", exp, res)
	}
}
