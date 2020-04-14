package P141

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	return race(head, head.Next)
}

func race(slow, fast *ListNode) bool {
	if slow == nil || fast == nil {
		return false
	}

	if slow == fast {
		return true
	}

	if fast.Next == nil {
		return false
	}

	return race(slow.Next, fast.Next.Next)
}
