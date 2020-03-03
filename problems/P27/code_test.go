package P27

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	exp := 5

	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	res := removeElement(nums, 2)

	fmt.Println(nums)

	if exp != res {
		t.Errorf("期待 %d 实际 %d\n", exp, res)
	}
}
