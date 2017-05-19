package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	want := []int{0, 1}
	got := twoSum([]int{2, 7, 11, 15}, 9)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}
