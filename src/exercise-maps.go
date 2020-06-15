package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	ret := make(map[string]int)
	fields := strings.Fields(s)
	for _, v := range fields {
		ret[v]++
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
