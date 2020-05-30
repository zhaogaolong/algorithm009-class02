package week01

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			break
		}
	}

	fast = head

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
