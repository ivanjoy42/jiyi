package main

import (
	"github.com/huichen/sego"
)

func freqWord(f []byte, std string) {
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("../text/dict.txt")
	se := segmenter.Segment(f)
	sego.SegmentsToSlice(se, false)
	text := sego.SegmentsToSlice(se, false)

	text = delChar(text)
	wc := count(text)
	wc = filter(wc, std)

	data := []WordCount{}
	for k, v := range wc {
		data = append(data, WordCount{k, v, 0, 0})
	}
	data = sortWord(data, 1)
	output(data, "../text/freqWord.txt")
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
