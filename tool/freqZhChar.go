package main

import (
	"gonum.org/v1/gonum/stat"
	"strings"
)

func freqChar(f, std []byte) {
	text := []rune(string(f))
	wc := count(text)
	standard := []rune(strings.ReplaceAll(string(std), "\n", ""))
	wc = filter(wc, standard)

	sd := dispersion(text)
	sd = filter(sd, standard)

	data := []WordCount{}
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

func dispersion(text []rune) map[rune]float64 {
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
