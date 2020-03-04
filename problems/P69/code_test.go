package P69

import (
	"testing"
)

func Test1(t *testing.T) {
	exp := 2
	res := mySqrt(4)

	if exp != res {
		t.Errorf("期待 %s 实际 %s\n", exp, res)
	}
}
