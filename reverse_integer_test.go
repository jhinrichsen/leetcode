package leetcode

import "testing"

const intSize = 32 + int(^uintptr(0)>>63<<5)

func TestReverseInteger0(t *testing.T) {
	want := 0
	got := reverse(0)
	if want != got {
		t.Fatalf("Expected %v but got %v", want, got)
	}
}

func TestReverseIntegerRegular(t *testing.T) {
	want := 321
	got := reverse(123)
	if want != got {
		t.Fatalf("Expected %v but got %v", want, got)
	}
}

func TestReverseIntegerNegative(t *testing.T) {
	want := -321
	got := reverse(-123)
	if want != got {
		t.Fatalf("Expected %v but got %v", want, got)
	}
}

func TestReverseIntegerOverflow(t *testing.T) {
	var willOverflow int
	switch intSize {
	case 32:
		// int32 max: 2147483647
		willOverflow = 1000000003
	case 64:
		// int64 max: 9223372036854775807
		willOverflow = 9223372036854775799
	}
	want := 0
	got := reverse(willOverflow)
	if want != got {
		t.Fatalf("Expected %v but got %v", want, got)
	}
}
