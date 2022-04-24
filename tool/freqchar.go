package main

import (
	"gonum.org/v1/gonum/stat"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

func freqChar(f []byte) {
	f2, _ := ioutil.ReadFile("../text/8105.txt")
	standard := regexp.MustCompile(`[\r\n]`).ReplaceAllString(string(f2), "")
	data := []WordCount{}
	text := string(f)
	wc := xCount(text)
	wc = filter(wc, standard)
	sd := dispersion(text)
	sd = filter(sd, standard)

	for k, v := range wc {
		data = append(data, WordCount{string(k), v, sd[k], 0})
	}
	data = combine(data)
	output(data, "../text/freqChar.txt")
}

func combine(data []WordCount) []WordCount {
	data = sortWord(data, 1)
	for i := range data {
		data[i].Order += i
	}

	data = sortWord(data, 2)
	for i := range data {
		data[i].Order += i
	}

	data = sortWord(data, 3)
	return data
}

func sortWord(data []WordCount, col int) []WordCount {
	sort.Slice(data, func(i, j int) bool {
		if col == 1 {
			return data[i].Count > data[j].Count
		} else if col == 2 {
			return int(data[i].StdDev) < int(data[j].StdDev)
		} else {
			return int(data[i].Order) < int(data[j].Order)
		}
	})
	return data
}

func xCount(text string) map[rune]int {
	wc := map[rune]int{}
	for _, v := range text {
		wc[v]++
	}
	return wc
}

func dispersion(text string) map[rune]float64 {
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
	return res
}

func filter[T any](wc map[rune]T, standard string) map[rune]T {
	for k := range wc {
		if !strings.ContainsRune(standard, k) {
			delete(wc, k)
		}
	}
	return wc
}
