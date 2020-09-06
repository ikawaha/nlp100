package chapter04

func VerbSurfaceFilter(ts []Token) []string {
	var ret []string
	for _, v := range ts {
		if v.POS == "動詞" {
			ret = append(ret, v.Surface)
		}
	}
	return ret
}
