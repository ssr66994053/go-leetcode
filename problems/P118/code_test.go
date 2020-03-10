package P118

import (
	"testing"
)

func Test1(t *testing.T) {
	exp := [][]int{
		[]int{1},
		[]int{1, 1},
		[]int{1, 2, 1},
		[]int{1, 3, 3, 1},
		[]int{1, 4, 6, 4, 1},
	}
	res := generate(5)

	if !equal(exp, res) {
		t.Errorf("期待 %v 实际 %v\n", exp, res)
	}
}

func equal(a1, a2 [][]int) bool {
	for i := 0; i < len(a1); i++ {
		for j := 0; j < len(a1[i]); j++ {
			if a1[i][j] != a2[i][j] {
				return false
			}
		}
	}

	return true
}
