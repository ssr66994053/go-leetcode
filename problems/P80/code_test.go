package P80

import (
	"fmt"
	"strconv"
	"testing"
)

func Test1(t *testing.T) {
	exp := build([]int{1, 2})
	head := build([]int{1, 1, 2})
	fmt.Println(toString(head))
	res := deleteDuplicates(head)

	if !equal(exp, res) {
		t.Errorf("期待 %s 实际 %s\n", toString(exp), toString(res))
	}
}

func build(arr []int) *ListNode {
	n := &ListNode{}
	h := n

	for i := 0; i < len(arr); i++ {
		if h.Val == 0 {
			h.Val = arr[i]
		} else {
			t := &ListNode{Val: arr[i]}
			h.Next = t
			h = h.Next
		}
	}

	return n
}

func equal(h1, h2 *ListNode) bool {
	if h1 != nil && h2 != nil {
		return h1.Val == h2.Val && equal(h1.Next, h2.Next)
	} else if h1 == nil && h2 == nil {
		return true
	}

	return false
}

func toString(h *ListNode) string {
	if h.Next == nil {
		return strconv.Itoa(h.Val)
	}

	return strconv.Itoa(h.Val) + "->" + toString(h.Next)
}
