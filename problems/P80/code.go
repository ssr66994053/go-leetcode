package P80

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	v := head.Val
	n := head.Next
	if n != nil {
		for v == n.Val {
			n = n.Next
			if n == nil {
				break
			}
		}

		head.Next = deleteDuplicates(n)
	}

	return head
}
