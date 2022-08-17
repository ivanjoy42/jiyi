package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func freqEnglish(f, scope []byte) (wf []WordFreq) {
	lemma, _ := ioutil.ReadFile(lemmaFile)
	text := splitWord(string(f))
	text = lemmatize(text, lemma)

	wc := count(text)
	standard := strings.Split(string(scope), "\n")
	wc = filter(wc, standard)

	for k, v := range wc {
		wf = append(wf, WordFreq{k, v, 0, 0, 0, 0})
	}
	wf = rank(wf)
	wf = freq(wf)

	return wf
}

func lemmatize(text []string, f []byte) []string {
	arr := strings.Split(string(f), "\n")
	lemma := map[string]string{}
	for _, v := range arr {
		tmp := strings.Split(v, "\t")
		lemma[tmp[1]] = tmp[0]
	}
	for i, v := range text {
		if _, ok := lemma[v]; ok {
			text[i] = lemma[v]
		}
	}
	return text
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
