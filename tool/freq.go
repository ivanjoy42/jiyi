package main

import (
	"fmt"
	"github.com/huichen/sego"
	"io/ioutil"
	"regexp"
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

	charText := charCount(string(f))
	charOut := sortChar(charText)
	ioutil.WriteFile("../text/freqchar.txt", []byte(charOut), 0644)

	f2, _ := ioutil.ReadFile("../text/booken.txt")
	entext := splitWord(string(f2))

	en := wordCount(entext)
	ioutil.WriteFile("../text/freqen.txt", en, 0644)
}

func charCount(text string) map[rune]int {
	wc := map[rune]int{}
	for _, v := range text {
		wc[v]++
	}

	f, _ := ioutil.ReadFile("../text/8105.txt")
	keys := regexp.MustCompile(`[\r\n]`).ReplaceAllString(string(f), "")
	for k := range wc {
		if !strings.ContainsRune(keys, k) {
			delete(wc, k)
		}
	}
	return wc
}

func sortChar(wc map[rune]int) (res string) {
	type WordCount struct {
		Word  rune
		Count int
	}

	ttl := 0
	a := []WordCount{}
	for k, v := range wc {
		a = append(a, WordCount{k, v})
		ttl += v
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].Count > a[j].Count
	})

	acc := 0
	for _, v := range a {
		acc += v.Count
		rate := float64(acc) / float64(ttl) * 100
		res += fmt.Sprintf("%s\t%d\t%f\n", string(v.Word), v.Count, rate)
	}
	return
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
		if line < 20000 {
			res += fmt.Sprintf("%s %d %f\n", v.Word, v.Count, rate)
			line++
		}
	}
	return []byte(res)
}
