package main

import (
	"strings"

	"gonum.org/v1/gonum/stat"
)

func freqChar(f, scope []byte) (wf []WordFreq) {
	text := []rune(string(f))

	wc := count(text)
	standard := []rune(strings.ReplaceAll(string(scope), "\n", ""))
	wc = filter(wc, standard)

	sd := dispersion(text)
	sd = filter(sd, standard)

	for k, v := range wc {
		wf = append(wf, WordFreq{string(k), v, 0, sd[k], 0, 0})
	}
	wf = rank(wf, len(standard))
	wf = freq(wf)

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
		pos[k] = append(pos[k], float64(ttl/cnt[k]))
	}

	res := map[rune]float64{}
	for k, v := range pos {
		res[k] = stat.StdDev(v, nil)
	}
	return res
}
