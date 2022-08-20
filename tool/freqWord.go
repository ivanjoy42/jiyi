package main

import (
	"strings"

	"github.com/huichen/sego"
)

func freqWord(f, scope []byte) (wf []WordFreq) {
	se := segmenter.Segment(f)
	sego.SegmentsToSlice(se, false)
	text := sego.SegmentsToSlice(se, false)

	wc := count(text)
	standard := strings.Split(string(scope), "\n")
	wc = filter(wc, standard)

	for k, v := range wc {
		wf = append(wf, WordFreq{k, v, 0, 0, 0, 0})
	}
	wf = rank(wf, len(standard))
	wf = freq(wf)

	return wf
}
