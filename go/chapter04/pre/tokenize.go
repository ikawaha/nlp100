package main

import (
	"bufio"
	"io"
	"strings"

	ipa "github.com/ikawaha/kagome-dict-ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
	"github.com/ikawaha/nlp100/go/chapter04"
)

type Tokenizer struct {
	tokenizer *tokenizer.Tokenizer
}

func NewTokenizer() (*Tokenizer, error) {
	t, err := tokenizer.New(ipa.Dict())
	return &Tokenizer{
		tokenizer: t,
	}, err
}
func (t Tokenizer) Tokenize(r io.Reader) ([]chapter04.Token, error) {
	var ret []chapter04.Token
	s := bufio.NewScanner(r)
	for s.Scan() {
		ret = append(ret, t.TokenizeSentence(strings.TrimSpace(s.Text()))...)
	}
	return ret, s.Err()
}

func (t Tokenizer) TokenizeSentence(s string) []chapter04.Token {
	ts := t.tokenizer.Tokenize(s)
	ret := make([]chapter04.Token, 0, len(ts))
	for _, v := range ts {
		surface := v.Surface
		base, _ := v.BaseForm()
		var (
			pos  string
			pos1 string
		)
		fs := v.POS()
		if len(fs) > 1 {
			pos = fs[0]
		}
		if len(fs) > 2 {
			pos1 = fs[1]
		}
		if v.Class == tokenizer.DUMMY {
			surface = ""
			pos = v.Surface
		}
		ret = append(ret, chapter04.Token{
			Surface: surface,
			Base:    base,
			POS:     pos,
			POS1:    pos1,
		})
	}
	return ret
}
