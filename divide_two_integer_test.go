package leetcode

import "testing"

func TestDivideTwoInteger(t *testing.T) {
	want := 3
	got := divide(21, 7)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}
