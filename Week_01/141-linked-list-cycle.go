package week01

import "fmt"

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil || head.Next.Next.Next == nil{
		return false
	}
	slow, fast := head, head.Next.Next
	coutner := 0
	for slow != fast {
		if fast == nil || fast.Next == nil || head.Next.Next == nil || head.Next.Next.Next == nil {
			return false
		}
		switch 5{
		case fast.Val, fast.Next.Val, fast.Next.Next.Val, fast.Next.Next.Next.Val:
			coutner++
			fmt.Printf("走了 %d 圈\n", coutner)
		}
		slow = slow.Next
		fast = fast.Next.Next.Next.Next.Next.Next.Next.Next
	}

	return true
}
