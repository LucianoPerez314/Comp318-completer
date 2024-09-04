package base

import (
	"sort"
	"strings"
)

type ByWord []string

func (s ByWord) Len() int           { return len(s) }
func (s ByWord) Less(i, j int) bool { return strings.Compare(s[i], s[j]) == -1 }
func (s ByWord) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type BaseObj struct {
	dict map[string]int
}

func (b BaseObj) Complete(base string) []string {
	//Extract valid keys from map
	var words = []string{}
	for word := range b.dict {
		if word == base {
			words = append(words, word)
		}
	}
	//Sort keys in alphabetical order
	sort.Sort(ByWord(words))
	//Return first 10 sorted keys
	length := len(words)
	if length >= 10 {
		return words[:10]
	} else {
		return words[:length]
	}

}

func New(mydict map[string]int) BaseObj {
	b := BaseObj{dict: mydict}
	return b
}
