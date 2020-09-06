package chapter04

import (
	"fmt"
	"testing"
)

func TestVerbSurfaceFilter(t *testing.T) {
	tokens, err := LoadTokenizedText("testdata/neko.txt.mecab")
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	base := VerbBaseFormFilter(tokens)
	fmt.Println(base)
}
