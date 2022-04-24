package main

import (
	"github.com/huichen/sego"
	"sort"
)

func freqWord(f []byte) {
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("../text/dict.txt")
	se := segmenter.Segment(f)
	sego.SegmentsToSlice(se, false)
	wordtext := sego.SegmentsToSlice(se, false)
	wordtext = delChar(wordtext)
	data := wordCount(wordtext)
	output(data, "../text/freqWord.txt")
}

func wordCount(text []string) []WordCount {
	wc := map[string]int{}
	for _, v := range text {
		if v != "\r" && v != "\n" {
			wc[v]++
		}
	}

	data := []WordCount{}
	for k, v := range wc {
		data = append(data, WordCount{k, v, 0, 0})
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].Count > data[j].Count
	})
	return data
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
