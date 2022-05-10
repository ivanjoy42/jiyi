package main

import (
	"fmt"
	"io/ioutil"
	"sort"
)

type WordCount struct {
	Word   string
	Count  int
	StdDev float64
	Order  int
}

func main() {
	bookCn, _ := ioutil.ReadFile("../text/book.txt")
	stdChar, _ := ioutil.ReadFile("../text/8105.txt")
	freqChar(bookCn, stdChar)

	stdWord, _ := ioutil.ReadFile("../text/13436.txt")
	freqWord(bookCn, stdWord)

	bookEn, _ := ioutil.ReadFile("../text/booken.txt")
	stdEnglish, _ := ioutil.ReadFile("../text/20000.txt")
	freqEnglish(bookEn, stdEnglish)
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
		if v.Count >= 5 && i < 20000 {
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
