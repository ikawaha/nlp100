package chapter03

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const jawikiCountryPath = "./testdata/jawiki-country.json.gz"

type Country struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func LoadTestdata() ([]Country, error) {
	f, err := os.Open(jawikiCountryPath)
	if err != nil {
		return nil, err
	}
	r, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	rd := bufio.NewReaderSize(r, 5*1024*1024)
	var ret []Country
	for eof := false; !eof; {
		line, err := rd.ReadBytes('\n')
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return ret, err
			}
			eof = true
		}
		if len(line) == 0 {
			continue
		}
		var c Country
		if err := json.Unmarshal(line, &c); err != nil {
			return ret, err
		}
		ret = append(ret, c)
	}
	return ret, nil
}

func Answer20() {
	cs, err := LoadTestdata()
	if err != nil {
		log.Fatalf("unexpected error %v", err)
	}
	for _, v := range cs {
		if v.Title == "イギリス" {
			fmt.Println(v.Text)
			break
		}
	}
}
