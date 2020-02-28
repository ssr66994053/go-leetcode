package P21

import (
	"strconv"
	"testing"
)

func Test1(t *testing.T) {
	exp := "1->1->2->3->4->4"
	l1 := build([]int{1, 2, 4})
	l2 := build([]int{1, 3, 4})
	res := flat(mergeTwoLists(l1, l2))
	if exp != res {
		t.Errorf("期待 %s 实际 %s\n", exp, res)
	}

	exp = "0"
	l1 = nil
	l2 = build([]int{0})
	res = flat(mergeTwoLists(l1, l2))
	if exp != res {
		t.Errorf("期待 %s 实际 %s\n", exp, res)
	}
}

func build(vals []int) *ListNode {
	var n, t *ListNode
	for i := 0; i < len(vals); i++ {
		if n == nil {
			n = &ListNode{}
			n.Val = vals[i]
			t = n
			continue
		}

		tmp := &ListNode{}
		tmp.Val = vals[i]
		t.Next = tmp
		t = t.Next
	}

	return n
}

func flat(n *ListNode) string {
	s := ""
	for n != nil {
		s = s + "->" + strconv.Itoa(n.Val)
		n = n.Next
	}

	return s[2:]
}
