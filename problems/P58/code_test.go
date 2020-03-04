package P58

import (
	"testing"
)

func Test1(t *testing.T) {
	exp := 5
	res := lengthOfLastWord("Hello World")

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}
