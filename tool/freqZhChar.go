package main

import (
	"gonum.org/v1/gonum/stat"
	"strings"
)

func freqChar(f []byte, std string) {
	data := []WordCount{}
	text := string(f)

	wc := count(text)
	wc = charFilter(wc, std)

	sd := dispersion(text)
	sd = charFilter(sd, std)

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

func count(text string) map[string]int {
	wc := map[string]int{}
	for _, v := range text {
		wc[string(v)]++
	}
	return wc
}

func dispersion(text string) map[string]float64 {
	last := map[string]int{}
	pos := map[string][]float64{}
	cnt := map[string]int{}

	for p, val := range text {
		v := string(val)
		pos[v] = append(pos[v], float64(p-last[v]))
		last[v] = p
		cnt[v]++
	}

	ttl := len(text)
	for k := range pos {
		pos[k] = append(pos[k], float64(ttl-last[k]))
		pos[k] = append(pos[k], float64(ttl/(cnt[k])))
	}

	res := map[string]float64{}
	for k, v := range pos {
		res[k] = stat.StdDev(v, nil)
	}
	return res
}

func charFilter[T any](wc map[string]T, std string) map[string]T {
	for k := range wc {
		if !strings.ContainsAny(string(k), std) {
			// if !strings.ContainsRune(std, k) {
			delete(wc, k)
		}
	}
	return wc
}
