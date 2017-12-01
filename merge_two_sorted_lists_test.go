package leetcode

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

// rand.Seed(time.Now().Unix())

// 2->4 + 3->5 => 2->3->4->5
func TestMergeTwoSortedLists(t *testing.T) {
	l1 := ListNode{2,
		&ListNode{4, nil},
	}
	l2 := ListNode{3,
		&ListNode{5, nil},
	}
	want := ListNode{2,
		&ListNode{3,
			&ListNode{4,
				&ListNode{5, nil},
			},
		},
	}
	got := mergeTwoLists(&l1, &l2)
	if cmpListNode(&want, got) != 0 {
		t.Fatalf("want %+v but got %+v\n", want, got)
	}
}

func TestMergeTwoSortedListsLimping(t *testing.T) {
	l1 := ListNode{2,
		&ListNode{4, nil},
	}
	l2 := ListNode{3, nil}
	want := ListNode{2,
		&ListNode{3,
			&ListNode{4, nil},
		},
	}
	got := mergeTwoLists(&l1, &l2)
	if cmpListNode(&want, got) != 0 {
		t.Fatalf("want %+v but got %+v\n", want, got)
	}
}

func cmpListNode(l1, l2 *ListNode) int {
	for {
		if l1 == nil && l2 == nil {
			return 0
		} else if l1 == nil && l2 != nil {
			return -1
		} else if l1 != nil && l2 == nil {
			return 1
		}

		if l1.Val < l2.Val {
			return -2
		} else if l1.Val > l2.Val {
			return 2
		}

		l1 = l1.Next
		l2 = l2.Next
	}
}

func TestCmp1(t *testing.T) {
	l1 := ListNode{2,
		&ListNode{4, nil},
	}
	l2 := ListNode{2,
		&ListNode{5, nil},
	}
	want := -2
	got := cmpListNode(&l1, &l2)
	if want != got {
		t.Fatalf("Want %d but got %d\n", want, got)
	}
}

func TestCmp2(t *testing.T) {
	l1 := ListNode{2,
		&ListNode{4, nil},
	}
	l2 := ListNode{2, nil}
	want := 1
	got := cmpListNode(&l1, &l2)
	if want != got {
		t.Fatalf("Want %d but got %d\n", want, got)
	}
}

func TestPermutations(t *testing.T) {
	rand.Seed(time.Now().Unix())
	l1 := newSortedListNode(t, rand.Intn(10))
	l2 := newSortedListNode(t, rand.Intn(10))
	want := l1.len() + l2.len()
	l := mergeTwoLists(l1, l2)
	got := l.len()
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func newSortedListNode(t *testing.T, n int) *ListNode {
	is := make([]int, n)
	for i := 0; i < n; i++ {
		is[i] = rand.Intn(10)
	}
	sort.Ints(is)
	return newListNode(t, len(is), func(i int) int {
		return is[i]
	})
}

func (a *ListNode) len() int {
	l := a
	n := 0
	for l != nil {
		l = l.Next
		n++
	}
	return n
}

func newListNode(t *testing.T, n int, generator func(int) int) *ListNode {
	var car *ListNode
	for i := n - 1; i >= 0; i-- {
		node := &ListNode{Val: generator(i), Next: car}
		car = node
	}
	// postcondition
	if car.len() != n {
		t.Fatalf("Want %v but got %v\n", n, car.len())
	}
	return car
}

func TestNilNil(t *testing.T) {
	got := mergeTwoLists(nil, nil)
	if got != nil {
		t.Fatalf("Want nil but got %v\n", got)
	}
}

func TestNilAny(t *testing.T) {
	want := &ListNode{0, nil}
	got := mergeTwoLists(nil, want)
	if cmpListNode(want, got) != 0 {
		t.Fatalf("Want %v but got %v\n", want, got)
	}
}
func TestAnyNil(t *testing.T) {
	want := &ListNode{0, nil}
	got := mergeTwoLists(want, nil)
	if cmpListNode(want, got) != 0 {
		t.Fatalf("Want %v but got %v\n", want, got)
	}
}
