package chapter04

import (
	"fmt"
	"log"
	"sort"

	"github.com/ikawaha/kagome/v2/tokenizer"
)

type Freq struct {
	Surface string
	Count   int
}

type FreqCounter map[string]int

func NewFreqCounter() *FreqCounter {
	return &FreqCounter{}
}

func (c FreqCounter) Add(ts ...tokenizer.Token) {
	for _, v := range ts {
		i := c[v.Surface]
		c[v.Surface] = i + 1
	}
}

func (c FreqCounter) List() []Freq {
	var list []Freq
	for k, v := range c {
		list = append(list, Freq{Surface: k, Count: v})
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i].Count == list[j].Count {
			return list[i].Surface < list[j].Surface
		}
		return list[i].Count > list[j].Count
	})
	return list
}

func Answer35() {
	sentences, err := TokenizeTextFile("./testdata/neko.txt")
	if err != nil {
		log.Fatalf("unexpected error, %v", err)
	}
	freq := NewFreqCounter()
	for _, s := range sentences {
		freq.Add(s...)
	}
	for i, v := range freq.List() {
		fmt.Println(i+1, v.Surface, v.Count)
	}
}
