package chapter04

import (
	"fmt"
	"log"
	"strings"

	"github.com/ikawaha/kagome/v2/filter"
)

func Answer34() {
	posFilter := filter.NewPOSFilter(filter.POS{"名詞"})
	sentences, err := TokenizeTextFile("./testdata/neko.txt")
	if err != nil {
		log.Fatalf("unexpected error, %v", err)
	}
	for _, s := range sentences {
		var phrase []string
		for i := 0; i < len(s); i++ {
			if posFilter.Match(s[i].POS()) {
				phrase = append(phrase, s[i].Surface)
				continue
			}
			if len(phrase) > 9 {
				fmt.Println(strings.Join(phrase, "/"))
			}
			phrase = phrase[0:0]
		}
	}
}
