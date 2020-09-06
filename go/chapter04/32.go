package chapter04

func VerbBaseFormFilter(ts []Token) []string {
	var ret []string
	for _, v := range ts {
		if v.POS == "動詞" {
			ret = append(ret, v.Base)
		}
	}
	return ret
}
