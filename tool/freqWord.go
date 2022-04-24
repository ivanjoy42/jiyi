package main

import (
	"github.com/huichen/sego"
	"strings"
)

func freqWord(f []byte, std string) {
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("../text/dict.txt")
	se := segmenter.Segment(f)
	sego.SegmentsToSlice(se, false)
	wordtext := sego.SegmentsToSlice(se, false)
	wordtext = delChar(wordtext)
	wc := wordCount(wordtext)
	wc = wordFilter(wc, std)

	data := []WordCount{}
	for k, v := range wc {
		data = append(data, WordCount{k, v, 0, 0})
	}
	data = sortWord(data, 1)
	output(data, "../text/freqWord.txt")
}

func wordCount(text []string) map[string]int {
	wc := map[string]int{}
	for _, v := range text {
		wc[v]++
	}
	return wc
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

func wordFilter(wc map[string]int, std string) map[string]int {
	for k := range wc {
		if !strings.ContainsAny(string(k), std) {
			delete(wc, k)
		}
	}
	return wc
}
