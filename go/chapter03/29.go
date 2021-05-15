package chapter03

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var (
	imageURLP = regexp.MustCompile(`"url":"(.+?)",`)
)

func ExtractCountryFlagFilename() ([]string, error) {
	data, err := ExtractBasicInformation()
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0, len(data))
	for _, info := range data {
		for k, v := range info.BasicInfo {
			if k != "国旗画像" {
				continue
			}
			p := strings.ReplaceAll(v, " ", "_")
			r, err := http.Get(`https://commons.wikimedia.org/w/api.php?action=query&prop=imageinfo&iiprop=url&format=json&titles=File:` + p)
			if err != nil {
				return nil, err
			}
			time.Sleep(time.Second)

			b, err := io.ReadAll(r.Body)
			if err != nil {
				return nil, err
			}
			match := imageURLP.FindStringSubmatch(string(b))
			if len(match) != 2 {
				continue
			}
			ret = append(ret, match[1])
			break
		}
	}
	return ret, nil
}

func Answer29() {
	urls, err := ExtractCountryFlagFilename()
	if err != nil {
		log.Printf("unexpected error, %v", err)
	}
	for _, v := range urls {
		fmt.Println(v)
	}
}
