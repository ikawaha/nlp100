package chapter04

import (
	"fmt"
	"testing"
)

func TestVerbBaseFormFilter(t *testing.T) {
	tokens, err := LoadTokenizedText("testdata/neko.txt.mecab")
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	base := VerbBaseFormFilter(tokens)
	fmt.Println(base)
}
