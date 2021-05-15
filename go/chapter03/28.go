package chapter03

import (
	"fmt"
	"log"
	"regexp"
)

var (
	htmlTagP = regexp.MustCompile(`<.+?>`)
)

func RemoveHTMLTag(s string) string {
	return htmlTagP.ReplaceAllString(s, "")
}

func RemoveMediaWikiMarkup(s string) string {
	ret := RemoveEmphasis(s)
	ret = RemoveMarkup(ret)
	ret = RemoveHTMLTag(ret)
	return ret
}

func ExtractBasicInformationWithoutMediaWikiMarkup() ([]BasicInfo, error) {
	info, err := ExtractBasicInformation()
	if err != nil {
		return nil, fmt.Errorf("extract basic information failed: %v", err)
	}
	for _, x := range info {
		fmt.Println(x.Country)
		for k, v := range x.BasicInfo {
			removed := RemoveMediaWikiMarkup(v)
			if v != removed {
				x.BasicInfo[k] = removed
			}
		}
	}
	return info, nil
}

func Answer28() {
	info, err := ExtractBasicInformationWithoutMediaWikiMarkup()
	if err != nil {
		log.Fatalf("unexpected error, %v", err)
	}
	for _, x := range info {
		fmt.Println(x.Country)
		for k, v := range x.BasicInfo {
			fmt.Println(k, "=", v)
		}
	}
}
