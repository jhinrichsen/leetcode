package leetcode

import (
	"log"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	samples := []int{3, 2, 2, 3}
	val := 3
	n := removeElement(samples, val)
	if n != 2 {
		log.Fatalf("Want 2 but got %v\n", n)
	}
	if samples[0] != 2 {
		log.Fatalf("Want index 0 to be 2 but got %v\n", samples[0])
	}
	if samples[1] != 2 {
		log.Fatalf("Want index 1 to be 2 but got %v\n", samples[1])
	}
}
