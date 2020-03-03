package P26

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	exp := 5

	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	res := removeDuplicates(nums)

	fmt.Println(nums)

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}
