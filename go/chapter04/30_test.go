package chapter04

import (
	"testing"
)

func TestLoadTokenizedText(t *testing.T) {
	tokens, err := LoadTokenizedText("testdata/neko.txt.mecab")
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if want, got := 206654, len(tokens); want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
