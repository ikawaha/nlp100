package chapter04

import (
	"fmt"
	"log"

	"github.com/ikawaha/kagome/v2/filter"
)

func Answer33() {
	posFilter := filter.NewPOSFilter(filter.POS{"名詞"})
	sentences, err := TokenizeTextFile("./testdata/neko.txt")
	if err != nil {
		log.Fatalf("unexpected error, %v", err)
	}
	for _, s := range sentences {
		for i := 2; i < len(s); i++ {
			if posFilter.Match(s[i-2].POS()) && s[i-1].Surface == "の" && posFilter.Match(s[i].POS()) {
				fmt.Println(s[i-2].Surface, s[i-1].Surface, s[i].Surface)
			}
		}
	}
}
