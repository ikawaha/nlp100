package chapter04

func FreqTop10Filter(ts []Token) []TokenFreq {
	list := Freq(ts)
	if len(list) < 10 {
		return list
	}
	return list[:10]
}
