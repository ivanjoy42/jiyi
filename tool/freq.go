package main

import (
	"fmt"
	"io/ioutil"
)

type WordCount struct {
	Word   string
	Count  int
	StdDev float64
	Order  int
}

func main() {
	f, _ := ioutil.ReadFile("../text/book.txt")
	freqChar(f)
	freqWord(f)

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
