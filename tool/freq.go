package main

import (
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/huichen/sego"
)

type WordFreq struct {
	Word     string
	Count    int
	Freq     float64
	Disperse float64
	Rank     int
	Score    float64
}

const (
	textDir = "../text/"
)

func main() {
	std := loadStd()

	batch(std, "char")
	batch(std, "word")
	batch(std, "english")
}

func loadStd() map[string][]byte {
	std := map[string][]byte{}
	std["char"], _ = ioutil.ReadFile(textDir + "8105.txt")
	std["word"], _ = ioutil.ReadFile(textDir + "13436.txt")
	std["english"], _ = ioutil.ReadFile(textDir + "20000.txt")
	return std
}

func batch(std map[string][]byte, stdType string) {
	stdDir := textDir + stdType + "/"
	files, _ := ioutil.ReadDir(stdDir)

	var segmenter sego.Segmenter
	if stdType == "word" {
		segmenter.LoadDictionary(textDir + "dict.txt")
	}

	scoreMap := map[string]float64{}
	for i, file := range files {
		println(stdType, i, file.Name())
		bookCn, _ := ioutil.ReadFile(stdDir + file.Name())

		wf := []WordFreq{}
		if stdType == "char" {
			wf = freqChar(bookCn, std[stdType])
		} else if stdType == "word" {
			wf = freqWord(bookCn, std[stdType], segmenter)
		} else if stdType == "english" {
			wf = freqEnglish(bookCn, std[stdType])
		}

		for _, v := range wf {
			scoreMap[v.Word] += v.Score
		}
	}

	score := []WordFreq{}
	for k, v := range scoreMap {
		score = append(score, WordFreq{k, 0, 0, 0, 0, v})
	}
	score = sortWord(score, 4, "desc")

	output(score, textDir+"freq"+stdType+".txt")
}

func output(wf []WordFreq, fileName string) {
	res := ""
	for _, v := range wf {
		res += fmt.Sprintf("%s\t%d\t%f\t%f\t%d\t%f\n", v.Word, v.Count, v.Freq, v.Disperse, v.Rank, v.Score)
	}
	ioutil.WriteFile(fileName, []byte(res), 0644)
}

func sortWord(wf []WordFreq, col int, order string) []WordFreq {
	sort.Slice(wf, func(i, j int) bool {
		if col == 1 {
			if order == "desc" {
				return wf[i].Count > wf[j].Count
			} else {
				return wf[i].Count < wf[j].Count
			}
		} else if col == 2 {
			if order == "desc" {
				return wf[i].Disperse > wf[j].Disperse
			} else {
				return wf[i].Disperse < wf[j].Disperse
			}
		} else if col == 3 {
			if order == "desc" {
				return wf[i].Rank > wf[j].Rank
			} else {
				return wf[i].Rank < wf[j].Rank
			}
		} else {
			if order == "desc" {
				return wf[i].Score > wf[j].Score
			} else {
				return wf[i].Score < wf[j].Score
			}
		}
	})
	return wf
}

func count[S rune | string](text []S) map[S]int {
	wc := map[S]int{}
	for _, v := range text {
		wc[v]++
	}
	return wc
}

func filter[N int | float64, S rune | string](wc map[S]N, std []S) map[S]N {
	intersect := map[S]int{}
	for _, v := range std {
		intersect[v]++
	}
	for k := range wc {
		intersect[k]++
	}
	for k, v := range intersect {
		if v < 2 {
			delete(wc, k)
		}
	}
	return wc
}
