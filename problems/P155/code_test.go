package P155

import (
	"testing"
)

func Test1(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	param_3 := obj.Top()
	if 3 != param_3 {
		t.Errorf("期待 %d 实际 %d\n", 3, param_3)
	}
	obj.Pop()
	param_1 := obj.GetMin()
	if 1 != param_1 {
		t.Errorf("期待 %d 实际 %d\n", 1, param_1)
	}
}
