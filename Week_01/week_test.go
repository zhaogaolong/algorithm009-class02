package week01

import (
	"testing"
)

func TestLink(t *testing.T) {

	root := &ListNode{
		Val: 0,
	}
	head := root
	var v *ListNode
	for i := 1; i <= 10; i++{
		root.Next = &ListNode{
			Val: i,
		}
		root = root.Next
		if i == 5{
			v = root
		}
		if i == 10{
			root.Next = v
		}
	}
	ok := hasCycle(head)
	if ok{
		t.Logf("yes")
	}

}

func TestRemoveDuplicate(t *testing.T) {
	data := []int{1, 1, 2}
	count := removeDuplicates(data)
	if count != 2 {
		t.Fatalf("want 2 got %d", count)
	}
}

func TestRemoveDuplicate2(t *testing.T) {
	data := []int{1, 1, 2}
	count := removeDuplicates(data)
	if count != 2 {
		t.Fail()
	}
	mergeTwoLists(nil, nil)
}
