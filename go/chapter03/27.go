package chapter03

import (
	"fmt"
	"log"
	"regexp"
)

var (
	markupP = regexp.MustCompile(`\[\[(?:[^|]*?\|)??([^|]*?)]]`)
)

func RemoveMarkup(s string) string {
	return markupP.ReplaceAllString(s, "${1}")
}

func ExtractBasicInformationWithoutEmphasisMarkup() ([]BasicInfo, error) {
	info, err := ExtractBasicInformationWithoutEmphasis()
	if err != nil {
		return nil, fmt.Errorf("extract basic information failed: %v", err)
	}
	for _, x := range info {
		fmt.Println(x.Country)
		for k, v := range x.BasicInfo {
			removed := RemoveMarkup(v)
			if v != removed {
				x.BasicInfo[k] = removed
			}
		}
	}
	return info, nil
}

func Answer27() {
	info, err := ExtractBasicInformationWithoutEmphasisMarkup()
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
