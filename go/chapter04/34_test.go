package chapter04

import (
	"fmt"
	"testing"
)

func TestLongestNounPhraseFilter(t *testing.T) {
	tokens, err := LoadTokenizedText("testdata/neko.txt.mecab")
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	np := LongestNounPhraseFilter(tokens)
	fmt.Println(np)
}
