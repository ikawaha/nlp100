package chapter04

import (
	"fmt"
	"log"

	"github.com/ikawaha/kagome/v2/filter"
)

func Answer36() {
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
		posFilter.Drop(&sentences[i]) // フィルターに含まれるものは落とす
		freq.Add(sentences[i]...)
	}
	list := freq.List()
	topK := 10
	for i := 0; i < topK && i < len(list); i++ {
		fmt.Println(i+1, list[i].Surface, list[i].Count)
	}
}
