package chapter04

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadTokenizedText(path string) ([]Token, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open file error, %v", err)
	}
	dec := json.NewDecoder(f)
	var ret []Token
	if err := dec.Decode(&ret); err != nil {
		return nil, fmt.Errorf("json decode error, %v", err)
	}
	return ret, nil
}
