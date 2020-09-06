package chapter04

import (
	"fmt"
	"testing"
)

func TestFreq(t *testing.T) {
	tokens, err := LoadTokenizedText("testdata/neko.txt.mecab")
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	list := Freq(tokens)
	for _, v := range list {
		fmt.Printf("surface:%s, base:%s, pos:%s, pos1:%s, freq:%d\n", v.token.Surface, v.token.Base, v.token.POS, v.token.POS1, v.freq)
	}
}
