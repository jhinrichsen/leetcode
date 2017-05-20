package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	want := []int{0, 1}
	got := TwoSum([]int{2, 7, 11, 15}, 9)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}

func (l1 *ListNode) Equal(l2 *ListNode) bool {
	for {
		// Value does not equal
		if (l1.Val != l2.Val) ||
			// One list ends, the other one does not
			(l1.Next == nil && l2.Next != nil) ||
			(l1.Next != nil && l2.Next == nil) {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}

func TestListNode(t *testing.T) {
	got := listNode(42)
	v := got.Val
	if v != 2 {
		t.Fatalf("want %v but got %v", 2, got)
	}
	v = got.Next.Val
	if v != 4 {
		t.Fatalf("want %v but got %v", 4, got)
	}
	if got.Next.Next != nil {
		t.Fatalf("want %v but got %v", nil, got)
	}
}

func TestVal(t *testing.T) {
	want := 2017
	got := val(listNode(want))
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
func TestAddTwoNumbers(t *testing.T) {
	l1 := listNode(342)
	l2 := listNode(465)

	want := 807
	got := val(AddTwoNumbers(l1, l2))
	if want != got {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}
