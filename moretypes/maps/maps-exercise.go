package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int) //map[key]answer
	for i := 0; i < len(words); i++ {
		value, ok := m[words[i]]
		if ok == false {
			m[words[i]] = value + 1
			continue
		}
		if ok != false && value > 0 {
			m[words[i]] = value + 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
