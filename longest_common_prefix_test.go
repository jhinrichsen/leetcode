package leetcode

import "testing"

func TestEvenEven(t *testing.T) {
	ss := []string{
		"abdc",
		"abec",
		"abfc",
	}
	want := "ab"
	got := longestCommonPrefix(ss)
	if want != got {
		t.Fatalf("Expected %v but got %v", want, got)
	}
}

func TestEvenOdd(t *testing.T) {
	ss := []string{
		"abcd",
		"abce",
		"abcf",
	}
	want := "abc"
	got := longestCommonPrefix(ss)
	if want != got {
		t.Fatalf("Expected %v but got %v", want, got)
	}
}

func TestOddEven(t *testing.T) {
	ss := []string{
		"abcde",
		"abcdf",
		"abcdg",
	}
	want := "abcd"
	got := longestCommonPrefix(ss)
	if want != got {
		t.Fatalf("Expected %v but got %v", want, got)
	}
}

func TestOddOdd(t *testing.T) {
	ss := []string{
		"abcdd",
		"abcee",
		"abcff",
	}
	want := "abc"
	got := longestCommonPrefix(ss)
	if want != got {
		t.Fatalf("Expected %v but got %v", want, got)
	}
}

func TestFox(t *testing.T) {
	ss := []string{
		"The quick brown fox jumps over the lazy cat",
		"The quick brown fox jumps over the lazy dog",
		"The quick brown fox jumps over the lazy donkey",
		"The quick brown fox jumps over the lazy mule",
		"The quick brown fox jumps over the lazy bird",
		"The quick brown fox jumps over the lazy whale",
		"The quick brown fox jumps over the lazy tiger",
		"The quick brown fox jumps over the lazy parrot",
	}
	want := "The quick brown fox jumps over the lazy "
	got := longestCommonPrefix(ss)
	if want != got {
		t.Fatalf("Expected >%v< but got >%v<", want, got)
	}
}

func TestAllEmpty(t *testing.T) {
	ss := []string{
		"",
		"",
		"",
		"",
	}
	want := ""
	got := longestCommonPrefix(ss)
	if want != got {
		t.Fatalf("Expected %v but got %v", want, got)
	}
}

func TestOneEmpty(t *testing.T) {
	ss := []string{
		"a",
		"b",
		"",
		"d",
	}
	want := ""
	got := longestCommonPrefix(ss)
	if want != got {
		t.Fatalf("Expected %v but got %v", want, got)
	}
}

// Windows, Intel Core i7-4770 @ 3.4 GHz: 428 ns/op
func BenchmarkBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ss := []string{
			"The quick brown fox jumps over the lazy cat",
			"The quick brown fox jumps over the lazy dog",
			"The quick brown fox jumps over the lazy donkey",
			"The quick brown fox jumps over the lazy mule",
			"The quick brown fox jumps over the lazy bird",
			"The quick brown fox jumps over the lazy whale",
			"The quick brown fox jumps over the lazy tiger",
			"The quick brown fox jumps over the lazy parrot",
		}
		want := "The quick brown fox jumps over the lazy "
		got := longestCommonPrefixBruteForce(ss)
		if want != got {
			b.Fatalf("Expected >%v< but got >%v<", want, got)
		}
	}
}

// Windows, Intel Core i7-4770 @ 3.4 GHz: 268 ns/op
func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ss := []string{
			"The quick brown fox jumps over the lazy cat",
			"The quick brown fox jumps over the lazy dog",
			"The quick brown fox jumps over the lazy donkey",
			"The quick brown fox jumps over the lazy mule",
			"The quick brown fox jumps over the lazy bird",
			"The quick brown fox jumps over the lazy whale",
			"The quick brown fox jumps over the lazy tiger",
			"The quick brown fox jumps over the lazy parrot",
		}
		want := "The quick brown fox jumps over the lazy "
		got := longestCommonPrefixBinarySearch(ss)
		if want != got {
			b.Fatalf("Expected >%v< but got >%v<", want, got)
		}
	}
}
