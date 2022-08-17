package main

import (
	"strings"

	"gonum.org/v1/gonum/stat"
)

func freqChar(f, std []byte) (wf []WordFreq) {
	text := []rune(string(f))

	wc := count(text)
	standard := []rune(strings.ReplaceAll(string(std), "\n", ""))
	wc = filter(wc, standard)

	sd := dispersion(text)
	sd = filter(sd, standard)

	for k, v := range wc {
		wf = append(wf, WordFreq{string(k), v, 0, sd[k], 0, 0})
	}
	wf = rank(wf)
	wf = freq(wf)

	return wf
}

func freq(wf []WordFreq) []WordFreq {
	ttl := 0
	for _, v := range wf {
		ttl += v.Count
	}

	for i, v := range wf {
		wf[i].Freq = float64(v.Count) / float64(ttl)
	}
	return wf
}

func rank(wf []WordFreq) []WordFreq {
	wf = sortWord(wf, 1, "desc")
	last := 0
	for i := range wf {
		if i > 0 && wf[i].Count == wf[i-1].Count {
			wf[i].Rank = last
		} else {
			wf[i].Rank = i
			last = i
		}
	}

	for i := range wf {
		wf[i].Score = 1 - float64(wf[i].Rank)/float64(last+1)
	}

	/*
		wf = sortWord(wf, 2)
		for i := range wf {
			wf[i].Order += i
		}
	*/

	wf = sortWord(wf, 3, "asc")
	return wf
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
