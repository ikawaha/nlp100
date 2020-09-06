package chapter04

import (
	"sort"
)

type Sentence []Token

func (s Sentence) hasNeko() bool {
	for _, v := range s {
		if v.Surface == "猫" {
			return true
		}
	}
	return false
}

func SentenceList(ts []Token) []Sentence {
	var ret []Sentence
	start := 0
	for i, v := range ts {
		if v.POS == "EOS" {
			ret = append(ret, ts[start:i])
			start = i + 1
		}
	}
	return ret
}

func CollocationNeko(ts []Token) []TokenFreq {
	ss := SentenceList(ts)
	freq := map[Token]int{}
	for _, s := range ss {
		if !s.hasNeko() {
			continue
		}
		for _, v := range s {
			if v.Surface == "猫" {
				continue
			}
			c := freq[v]
			freq[v] = c + 1
		}
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
	if len(list) < 10 {
		return list
	}
	return list[:10]
}
