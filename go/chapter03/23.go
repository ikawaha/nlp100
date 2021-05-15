package chapter03

import (
	"fmt"
	"log"
	"regexp"
)

var (
	sectionP = regexp.MustCompile(`(?m)^(==+)\s*([^=]+?)\s*(==+)$`)
)

type Section struct {
	Name  string
	Level int
}

func ExtractSection() ([]Section, error) {
	data, err := LoadTestdata()
	if err != nil {
		log.Fatalf("unexpected error, %v", err)
	}
	var ret []Section
	for _, v := range data {
		matches := sectionP.FindAllStringSubmatch(v.Text, -1)
		for _, match := range matches {
			if len(match) < 4 {
				continue
			}
			if match[1] != match[3] {
				continue
			}
			ret = append(ret, Section{
				Name:  match[2],
				Level: len(match[1]) - 1,
			})
		}
	}
	return ret, nil
}

func Answer23() {
	cs, err := ExtractSection()
	if err != nil {
		log.Fatalf("unexpected error, %v", err)
	}
	for _, v := range cs {
		fmt.Println(v.Name, v.Level)
	}
}
