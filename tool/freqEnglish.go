package main

import (
	"regexp"
	"strings"
)

func freqEnglish(f []byte) {
	text := splitWord(string(f))
	wc := count(text)
	data := []WordCount{}
	for k, v := range wc {
		data = append(data, WordCount{k, v, 0, 0})
	}
	data = sortWord(data, 1)

	output(data, "../text/freqEnglish.txt")
}

func splitWord(text string) []string {
	text = strings.ReplaceAll(text, "’", "'")
	text = strings.ReplaceAll(text, "—", "-")
	arr := regexp.MustCompile(`[^a-zA-Z'-]+`).Split(text, -1)
	res := []string{}
	for _, v := range arr {
		if v != "" && !strings.ContainsAny(v, "'-ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			res = append(res, v)
		}
	}
	return res
}
