package frequency

import (
	"sort"
	"strings"
)

type Entry struct {
	word      string
	frequency int
}

type ByFrequency []Entry

func (e ByFrequency) Len() int           { return len(e) }
func (e ByFrequency) Less(i, j int) bool { return e[i].frequency < e[j].frequency }
func (e ByFrequency) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

type FrequencyObj struct {
	dict map[string]int
}

func (f FrequencyObj) Complete(base string) []string {
	//Extract valid entries from map
	entries := []Entry{}
	for key := range f.dict {
		if strings.HasPrefix(key, base) {
			entry := Entry{word: key, frequency: f.dict[key]}
			entries = append(entries, entry)
		}
	}
	//Sort by frequency.
	sort.Sort(ByFrequency(entries))
	//Extract strings from entries
	length := len(entries)
	words := []string{}
	for i := 0; i < length; i++ {
		words = append(words, entries[i].word)
	}
	//Return first 10 words

	if length >= 10 {
		return words[:10]
	} else {
		return words[:length]
	}
}

func New(mydict map[string]int) FrequencyObj {
	b := FrequencyObj{dict: mydict}
	return b
}
