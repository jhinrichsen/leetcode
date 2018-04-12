package leetcode

import (
	"log"
	"testing"
)

func TestSample(t *testing.T) {
	in := []int{1, 1, 2}
	n := RemoveDuplicatesFromSortedArray(in)
	if n != 2 {
		log.Fatalf("Want 2 but got %v\n", n)
	}
	if in[0] != 1 {
		log.Fatalf("Want index 0 to be 1 but got %v\n", in[0])
	}
	if in[1] != 2 {
		log.Fatalf("Want index 0 to be 1 but got %v\n", in[1])
	}
}
