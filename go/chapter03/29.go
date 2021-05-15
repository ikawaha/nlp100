package chapter03

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var (
	imageURLP    = regexp.MustCompile(`"url":"(.+?)",`)
	errNotFound  = errors.New("not found")
	wikiMediaURL = `https://commons.wikimedia.org/w/api.php?action=query&prop=imageinfo&iiprop=url&format=json&titles=File:`
)

func fetch(p string) (string, error) {
	r, err := (&http.Client{
		Timeout: 30 * time.Second,
	}).Get(wikiMediaURL + p)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	match := imageURLP.FindStringSubmatch(string(b))
	if len(match) == 2 {
		return match[1], nil
	}
	return "", errNotFound
}

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
			url, err := fetch(p)
			if err != nil {
				if errors.Is(err, errNotFound) {
					continue
				}
				return nil, err
			}
			time.Sleep(time.Second)
			ret = append(ret, url)
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
