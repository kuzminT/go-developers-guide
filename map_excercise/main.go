package main

import (
	// "golang.org/x/tour/wc"
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	wordsSlice := strings.Fields(s)

	wordsCountMap := make(map[string]int)

	for _, word := range wordsSlice {
		if _, ok := wordsCountMap[word]; ok {
			wordsCountMap[word] += 1
		} else {
			wordsCountMap[word] = 1
		}
	}

	return wordsCountMap
}

func main() {
	// wc.Test(WordCount)
	fmt.Print(WordCount("I ate a donut. Then I ate another donut."))
}
