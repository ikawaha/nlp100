package chapter04

import (
	"fmt"
	"log"

	"github.com/ikawaha/kagome/v2/filter"
)

func Answer31() {
	posFilter := filter.NewPOSFilter(filter.POS{"動詞"}) // 動詞だけがセットされた品詞フィルター
	sentences, err := TokenizeTextFile("./testdata/neko.txt")
	if err != nil {
		log.Fatalf("unexpected error, %v", err)
	}
	for _, s := range sentences {
		for _, token := range s {
			if posFilter.Match(token.POS()) {
				fmt.Println(token.Surface)
			}
		}
	}
}
