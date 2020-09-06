package chapter04

import (
	"fmt"
	"testing"
)

func TestNounPhraseANOBFilter(t *testing.T) {
	tokens, err := LoadTokenizedText("testdata/neko.txt.mecab")
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	nps := NounPhraseANOBFilter(tokens)
	fmt.Println(nps)
}
