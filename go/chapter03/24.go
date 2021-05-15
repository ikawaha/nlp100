package chapter03

import (
	"fmt"
	"log"
	"regexp"
)

var (
	fileP = regexp.MustCompile(`(?m)\[\[ファイル:([^|\]]+)`)
)

func ExtractFile() ([]string, error) {
	data, err := LoadTestdata()
	if err != nil {
		log.Fatalf("unexpected error, %v", err)
	}
	var ret []string
	for _, v := range data {
		matches := fileP.FindAllStringSubmatch(v.Text, -1)
		for _, match := range matches {
			if len(match) < 1 {
				continue
			}
			ret = append(ret, match[1])
		}
	}
	return ret, nil
}

func Answer24() {
	cs, err := ExtractFile()
	if err != nil {
		log.Fatalf("unexpected error, %v", err)
	}
	for _, v := range cs {
		fmt.Println(v)
	}
}
