package chapter04

import (
	"sort"
)

type TokenFreq struct {
	token Token
	freq  int
}

func Freq(ts []Token) []TokenFreq {
	freq := map[Token]int{}
	for _, v := range ts {
		if v.POS == "BOS" || v.POS == "EOS" {
			continue
		}
		c := freq[v]
		freq[v] = c + 1
	}
	list := make([]TokenFreq, 0, len(freq))
	for k, v := range freq {
		list = append(list, TokenFreq{
			token: k,
			freq:  v,
		})
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i].freq == list[j].freq {
			return list[i].token.Surface < list[j].token.Surface
		}
		return list[i].freq > list[j].freq
	})
	return list
}
