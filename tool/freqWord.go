package main

import (
	"github.com/huichen/sego"
	"strings"
)

func freqWord(f, std []byte) {
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("../text/dict.txt")
	se := segmenter.Segment(f)
	sego.SegmentsToSlice(se, false)

	text := sego.SegmentsToSlice(se, false)
	wc := count(text)
	standard := strings.Split(string(std), "\n")
	wc = filter(wc, standard)

	data := []WordCount{}
	for k, v := range wc {
		data = append(data, WordCount{k, v, 0, 0})
	}
	data = sortWord(data, 1)
	output(data, "../text/freqWord.txt")
}
