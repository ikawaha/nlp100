package chapter03

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

var (
	basicInfoP = regexp.MustCompile(`(?ms)\{\{基礎情報(.*?)^(?:\}\}|''')`) // 基礎情報の末尾が `^}}` で取れないことがある
	infoP      = regexp.MustCompile(`\s*\|\s*([^=]+?)\s*=\s*(.*)$`)
)

type BasicInfo struct {
	Country   string
	BasicInfo map[string]string
}

func ExtractBasicInformation() ([]BasicInfo, error) {
	data, err := LoadTestdata()
	if err != nil {
		log.Fatalf("unexpected error, %v", err)
	}
	var ret []BasicInfo
	for _, v := range data {
		matches := basicInfoP.FindAllStringSubmatch(v.Text, -1)
		for _, match := range matches {
			if len(match) != 2 {
				continue
			}
			r := BasicInfo{
				Country:   v.Title,
				BasicInfo: map[string]string{},
			}
			lines := strings.Split(match[1], "\n")
			for _, line := range lines {
				v := infoP.FindStringSubmatch(line)
				if len(v) < 3 {
					continue
				}
				r.BasicInfo[v[1]] = v[2]
			}
			ret = append(ret, r)
		}
	}
	return ret, nil
}

func Answer25() {
	info, err := ExtractBasicInformation()
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
