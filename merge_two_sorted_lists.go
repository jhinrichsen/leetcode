package leetcode

/* Same structure as add_two_numbers
type ListNode struct {
	Val  int
	Next *ListNode
}
*/

// Merge two sorted linked lists and return it as a new list. The new list
// should be made by splicing together the nodes of the first two lists.
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	min := func(l1, l2 *ListNode) *ListNode {
		if l1 == nil {
			return l2
		}
		if l2 == nil {
			return l1
		}
		if l2.Val < l1.Val {
			return l2
		}
		return l1
	}

	car := min(l1, l2)
	cons := car

	left, right := l1, l2
	for {
		if cons == left {
			left = left.Next
		} else {
			right = right.Next
		}
		l := min(left, right)
		if l == nil {
			break
		}
		cons.Next = l
		cons = cons.Next
	}
	return car
}
