package main

import (
	"strings"

	"github.com/huichen/sego"
)

func freqWord(f, std []byte, segmenter sego.Segmenter) (wf []WordFreq) {
	se := segmenter.Segment(f)
	sego.SegmentsToSlice(se, false)
	text := sego.SegmentsToSlice(se, false)

	wc := count(text)
	standard := strings.Split(string(std), "\n")
	wc = filter(wc, standard)

	for k, v := range wc {
		wf = append(wf, WordFreq{k, v, 0, 0, 0, 0})
	}
	wf = rank(wf)

	return wf
}
