package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

func freqChar() {
	f, _ := ioutil.ReadFile("../text/zrx.txt")
	
	f2, _ := ioutil.ReadFile("../text/8105.txt")
	standard := regexp.MustCompile(`[\r\n]`).ReplaceAllString(string(f2), "")

	charText := charCount(string(f), standard)
	charOut := sortChar(charText)
	ioutil.WriteFile("../text/freqchar.txt", []byte(charOut), 0644)

	d := dispersion(string(f), standard)
	dOut := sortChar2(d)
	ioutil.WriteFile("../text/freqchar2.txt", []byte(dOut), 0644)
}

func dispersion(text, standard string) map[rune]float64 {
	last := map[rune]int{}
	pos := map[rune][]float64{}
	cnt := map[rune]int{}

	for p, v := range text {
		pos[v] = append(pos[v], float64(p-last[v]))
		last[v] = p
		cnt[v]++
	}

	ttl := len(text)
	for k := range pos {
		pos[k] = append(pos[k], float64(ttl-last[k]))
		pos[k] = append(pos[k], float64(ttl/(cnt[k])))
	}

	res := map[rune]float64{}
	for k, v := range pos {
		res[k] = stat.StdDev(v, nil)
	}

	for k := range res {
		if !strings.ContainsRune(standard, k) {
			delete(res, k)
		}
	}

	return res
}

func charCount(text, standard string) map[rune]int {
	wc := map[rune]int{}
	for _, v := range text {
		wc[v]++
	}

	for k := range wc {
		if !strings.ContainsRune(standard, k) {
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
		// rate := float64(acc) / float64(ttl) * 100
		res += fmt.Sprintf("%s\t%d\n", string(v.Word), v.Count)
		// res += fmt.Sprintf("%s\t%d\t%f\n", string(v.Word), v.Count, rate)
	}
	return
}

func sortChar2(wc map[rune]float64) (res string) {
	type WordCount struct {
		Word  rune
		Count float64
	}

	ttl := 0.0
	a := []WordCount{}
	for k, v := range wc {
		a = append(a, WordCount{k, v})
		ttl += v
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].Count < a[j].Count
	})

	acc := 0.0
	for _, v := range a {
		acc += v.Count
		// rate := float64(acc) / float64(ttl) * 100
		res += fmt.Sprintf("%s\t%f\n", string(v.Word), v.Count)
		// res += fmt.Sprintf("%s\t%f\t%f\n", string(v.Word), v.Count, rate)
	}
	return
}
