package chapter04

func LongestNounPhraseFilter(ts []Token) []string {
	maxRnage := struct {
		start int
		end   int
	}{}
	var (
		inside bool
		start  int
	)
	for i := range ts {
		if ts[i].POS == "名詞" {
			if !inside {
				inside = true
				start = i
			}
		} else if inside {
			inside = false
			if maxRnage.end-maxRnage.start < i-start {
				maxRnage.start = start
				maxRnage.end = i
			}
		}
	}
	ret := make([]string, 0, maxRnage.end-maxRnage.start)
	for i := maxRnage.start; i < maxRnage.end; i++ {
		ret = append(ret, ts[i].Surface)
	}
	return ret
}
