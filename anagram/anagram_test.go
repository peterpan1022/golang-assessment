package anagram

import (
	"log"
	"testing"
)

func TestCheckAnagram(t *testing.T) {
	word1 := "cat"
	word2 := "act"

	result := CheckAnagram(word1, word2)

	if !result.IsAnagram {
		t.Fatalf(`checkAnagram("cat", "act").IsAnagram: = %v, want %v`, result.IsAnagram, true)
	}

	if result.Distance != 2 {
		t.Fatalf(`checkAnagram("cat", "act").Distance: = %v, want %v`, result.Distance, 2)
	}
}
