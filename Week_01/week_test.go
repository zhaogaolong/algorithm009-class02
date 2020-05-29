package week01

import (
	"testing"
)

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
