package problem1

import "testing"

func Test1(t *testing.T) {
	exp := []int{0, 1}
	res := twoSum([]int{2, 7, 11, 13}, 9)
	for i, n := range res {
		if exp[i] != n {
			t.Errorf("期待 %d 实际 %d\n", exp[i], n)
		}
	}
}

func Test2(t *testing.T) {
	exp := []int{0, 3}
	res := twoSum([]int{0, 3, 4, 0}, 0)
	for i, n := range res {
		if exp[i] != n {
			t.Errorf("期待 %d 实际 %d\n", exp[i], n)
		}
	}
}
