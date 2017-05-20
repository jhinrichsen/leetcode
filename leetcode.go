package leetcode

func TwoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j, w := range nums {
			if i == j {
				continue
			}
			if v+w == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func val(l *ListNode) int {
	n := 0
	exp := 1
	for {
		if l == nil {
			break
		}
		n += l.Val * exp
		exp *= 10
		l = l.Next
	}
	return n
}

func listNode(i int) *ListNode {
	head := &ListNode{}
	l := head
	for {
		l.Val = i % 10
		i /= 10
		if i == 0 {
			break
		}
		l.Next = &ListNode{}
		l = l.Next
	}
	return head
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	i1 := val(l1)
	i2 := val(l2)

	return listNode(i1 + i2)
}
