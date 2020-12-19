package chapter04

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

// 1文毎に形態素解析した結果。型変換がめんどいので alias
type TokenizedSentence = []tokenizer.Token

type Tokenizer struct {
	*tokenizer.Tokenizer
}

func NewTokenizer() (*Tokenizer, error) {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos()) // BOS/EOS は結果に含めない
	return &Tokenizer{
		Tokenizer: t,
	}, err
}

// Reader から1行毎にテキストを読み取って、形態素解析をおこなう
func (t Tokenizer) TokenizeReader(r io.Reader) ([]TokenizedSentence, error) {
	s := bufio.NewScanner(r)
	var ret []TokenizedSentence
	for s.Scan() {
		sen := strings.TrimSpace(s.Text())         // 一行一文なので1行を取り出す
		ret = append(ret, t.TokenizeSentence(sen)) // 形態素解析した結果を詰める
	}
	return ret, s.Err()
}

// 1文の単位を形態素する
func (t Tokenizer) TokenizeSentence(s string) TokenizedSentence {
	return t.Tokenize(s)
}

// ファイルパスを与えると1行毎に形態素解析して返す
func TokenizeTextFile(path string) ([]TokenizedSentence, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open file error, %v", err)
	}
	t, err := NewTokenizer()
	if err != nil {
		return nil, err
	}
	return t.TokenizeReader(f)
}

func Answer30() {
	sentences, err := TokenizeTextFile("./testdata/neko.txt")
	if err != nil {
		log.Fatalf("unexpected error, %v", err)
	}
	for i, v := range sentences {
		fmt.Println(i, v)
	}
}
