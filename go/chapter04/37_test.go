package chapter04

import (
	"fmt"
	"testing"
)

func TestCollocationNeko(t *testing.T) {
	tokens, err := LoadTokenizedText("testdata/neko.txt.mecab")
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	list := CollocationNeko(tokens)
	if want, got := 10, len(list); want != got {
		t.Errorf("want %v, got %v", want, got)
	}
	for _, v := range list {
		fmt.Printf("surface:%s, base:%s, pos:%s, pos1:%s, freq:%d\n", v.token.Surface, v.token.Base, v.token.POS, v.token.POS1, v.freq)
	}
}
