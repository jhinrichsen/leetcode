package leetcode

// Given an array of integers, return indices of the two numbers such that they
// add up to a specific target.
// You may assume that each input would have exactly one solution, and you may
// not use the same element twice.
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

/*
type ListNode struct {
	Val  int
	Next *ListNode
}

func val(l *ListNode) *big.Int {
	n := big.NewInt(0)
	exp := big.NewInt(1)
	ten := big.NewInt(10)
	for {
		if l == nil {
			break
		}
		// n += l.Val * exp
		v := big.NewInt(int64(l.Val))
		v.Mul(v, exp)
		n.Add(n, v)
		exp.Mul(exp, ten)
		l = l.Next
	}
	return n
}

func listNode(i *big.Int) *ListNode {
	head := &ListNode{}
	l := head
	zero := big.NewInt(0)
	ten := big.NewInt(10)
	v := big.NewInt(0)
	for {
		// last digit can always be downcasted into an int
		l.Val = int(v.Mod(i, ten).Int64())
		i.Div(i, ten)
		if i.Cmp(zero) == 0 {
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

	return listNode(i1.Add(i1, i2))
}

// LengthOfLongestSubstring returns number of bytes (not runes)
func LengthOfLongestSubstring(s string) int {
	max := 0
	for j := 0; j < len(s); j++ {
		m := make(map[byte]bool)
		for i := j; i < len(s); i++ {
			c := s[i]
			if m[c] {
				break
			}
			m[c] = true
		}
		if len(m) > max {
			max = len(m)
		}
	}
	return max
}
*/
