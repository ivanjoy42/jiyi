package main

import (
	"fmt"
	"github.com/huichen/sego"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	f, _ := ioutil.ReadFile("../text/book.txt")

	var segmenter sego.Segmenter
	segmenter.LoadDictionary("../text/dict.txt")

	se := segmenter.Segment(f)
	wordtext := sego.SegmentsToSlice(se, false)
	word := wordCount(delChar(wordtext))
	ioutil.WriteFile("../text/freqword.txt", word, 0644)

	chartext := strings.Split(string(f), "")
	char := wordCount(chartext)
	ioutil.WriteFile("../text/freqchar.txt", char, 0644)
}

func delChar(text []string) []string {
	res := []string{}
	for _, v := range text {
		if len(v) >= 6 {
			res = append(res, v)
		}
	}
	return res
}

func wordCount(text []string) []byte {
	res := ""
	total := 0
	wc := map[string]int{}
	for _, v := range text {
		if v != "\r" && v != "\n" {
			wc[v]++
			total++
		}
	}

	type WordCount struct {
		Word  string
		Count int
	}

	a := []WordCount{}
	for k, v := range wc {
		a = append(a, WordCount{k, v})
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].Count > a[j].Count
	})

	line := 0
	acc := 0
	for _, v := range a {
		acc += v.Count
		rate := float64(acc) / float64(total) * 100
		//if v.Count >= 1 && rate <= 99 && line <= 20000 {
		if line < 20000 {
			res += fmt.Sprintf("%s %d %f\n", v.Word, v.Count, rate)
			line++
		}
	}
	return []byte(res)
}
