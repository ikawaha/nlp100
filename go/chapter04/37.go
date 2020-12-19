package chapter04

import (
	"fmt"
	"log"

	"github.com/ikawaha/kagome/v2/filter"
)

func Answer37() {
	wordFilter := filter.NewWordFilter([]string{"猫"})
	posFilter := filter.NewPOSFilter([]filter.POS{
		{"助詞"},              // が も の を
		{"助動詞"},             // れる られる せる
		{"記号"},              // 。、？
		{"名詞", "非自立", "一般"}, // の
	}...)
	sentences, err := TokenizeTextFile("./testdata/neko.txt")
	if err != nil {
		log.Fatalf("unexpected error, %v", err)
	}
	freq := NewFreqCounter()
	for i := range sentences {
		n := len(sentences[i])         // 元の文の長さを記録しておいて
		wordFilter.Drop(&sentences[i]) // 猫が入っていたらその形態素を落とす
		if n != len(sentences[i]) {    // 長さが変わっていたら共起している形態素を足す
			posFilter.Drop(&sentences[i]) // 面白みのない形態素を落とす
			freq.Add(sentences[i]...)
		}
	}
	for i, v := range freq.List() {
		fmt.Println(i+1, v.Surface, v.Count)
	}
}
