package chapter04

type NounPhraseANOB struct {
	A string
	B string
}

func (n NounPhraseANOB) String() string {
	return n.A + "の" + n.B
}

func NounPhraseANOBFilter(ts []Token) []NounPhraseANOB {
	var ret []NounPhraseANOB
	for i := range ts {
		if i < 2 {
			continue
		}
		if ts[i-2].POS == "名詞" && ts[i-1].Surface == "の" && ts[i].POS == "名詞" {
			ret = append(ret, NounPhraseANOB{
				A: ts[i-2].Surface,
				B: ts[i].Surface,
			})
		}
	}
	return ret
}
