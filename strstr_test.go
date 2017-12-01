package leetcode

import (
	"log"
	"testing"
)

func TestStrstrFound(t *testing.T) {
	haystack := "hello"
	needle := "ll"
	want := 2
	got := strstr(haystack, needle)
	if want != got {
		log.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestStrstrNotFound(t *testing.T) {
	haystack := "aaaaa"
	needle := "bba"
	want := -1
	got := strstr(haystack, needle)
	if want != got {
		log.Fatalf("want %v but got %v\n", want, got)
	}
}
