package chapter03

import (
	"fmt"
	"log"
	"regexp"
)

var (
	emphasisP = regexp.MustCompile(`('{2,5})(.*?)('{2,5})`)
)

func RemoveEmphasis(s string) string {
	idx := emphasisP.FindAllStringIndex(s, -1)
	seg := emphasisP.FindAllStringSubmatch(s, -1)
	ret := s
	for i := len(idx) - 1; i >= 0; i-- {
		if len(seg[i]) != 4 {
			continue
		}
		if seg[i][1] != seg[i][3] {
			continue
		}
		ret = ret[:idx[i][0]] + seg[i][2] + ret[idx[i][1]:]

	}
	return ret
}

func ExtractBasicInformationWithoutEmphasis() ([]BasicInfo, error) {
	info, err := ExtractBasicInformation()
	if err != nil {
		return nil, fmt.Errorf("extract basic information failed: %v", err)
	}
	for _, x := range info {
		fmt.Println(x.Country)
		for k, v := range x.BasicInfo {
			removed := RemoveEmphasis(v)
			if v != removed {
				x.BasicInfo[k] = removed
			}
		}
	}
	return info, nil
}

func Answer26() {
	info, err := ExtractBasicInformationWithoutEmphasis()
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
