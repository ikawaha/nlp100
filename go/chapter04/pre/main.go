package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	commandName = "nlp100-go-chapter04-pre"
)

// 夏目漱石の小説『吾輩は猫である』の文章（neko.txt）をMeCabを使って形態素解析し，
// その結果をneko.txt.mecabというファイルに保存せよ
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	opt := NewOption()
	if err := opt.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("option parse error, %v", err)
	}
	var in *os.File = os.Stdin
	if opt.Input != "" {
		var err error
		in, err = os.Open(opt.Input)
		if err != nil {
			return fmt.Errorf("cannot open input file, %v", err)
		}
		defer in.Close()
	}
	tokenizer, err := NewTokenizer()
	if err != nil {
		return fmt.Errorf("new tokenizer error, %v", err)
	}
	tokens, err := tokenizer.Tokenize(in)
	if err != nil {
		fmt.Errorf("tokenize error, %v", err)
	}
	b, err := json.Marshal(tokens)
	if err != nil {
		return fmt.Errorf("data convert error, %v", err)
	}
	var out *os.File = os.Stdout
	if opt.Output != "" {
		var err error
		out, err = os.Create(opt.Output)
		if err != nil {
			return fmt.Errorf("cannot open output file, %v", err)
		}
		defer out.Close()
	}

	if _, err := out.Write(b); err != nil {
		return fmt.Errorf("write error, %v", err)
	}
	return nil
}
