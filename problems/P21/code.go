package P21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var h, t *ListNode
	if l1.Val <= l2.Val {
		h = l1
		t = l1
		l1 = l1.Next
	} else {
		h = l2
		t = l2
		l2 = l2.Next
	}

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			t.Next = l1
			l1 = l1.Next
		} else {
			t.Next = l2
			l2 = l2.Next
		}
		t = t.Next
	}

	if l1 != nil {
		t.Next = l1
	}
	if l2 != nil {
		t.Next = l2
	}

	return h
}
