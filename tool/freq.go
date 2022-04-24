package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

type WordCount struct {
	Word   string
	Count  int
	StdDev float64
	Order  int
}

func main() {
	f, _ := ioutil.ReadFile("../text/8105.txt")
	std := regexp.MustCompile(`[\r\n]`).ReplaceAllString(string(f), "")

	f1, _ := ioutil.ReadFile("../text/book.txt")
	freqChar(f1, std)
	freqWord(f1, std)

	f2, _ := ioutil.ReadFile("../text/booken.txt")
	freqEnglish(f2)
}

func output(data []WordCount, fileName string) {
	ttl := 0
	for _, v := range data {
		ttl += v.Count
	}

	res := ""
	acc := 0
	for i, v := range data {
		acc += v.Count
		rate := float64(acc) / float64(ttl) * 100
		if i < 20000 {
			res += fmt.Sprintf("%s\t%d\t%.2f%%\t%.0f\n", v.Word, v.Count, rate, v.StdDev)
		}
	}
	ioutil.WriteFile(fileName, []byte(res), 0644)
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

func count[S rune | string](text []S) map[S]int {
	wc := map[S]int{}
	for _, v := range text {
		wc[v]++
	}
	return wc
}

func filter[N any, S rune | string](wc map[S]N, std string) map[S]N {
	for k := range wc {
		if !strings.ContainsAny(string(k), std) {
			delete(wc, k)
		}
	}
	return wc
}
