package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	dict := make(map[string]int)
	for _, v := range strings.Fields(s) {
		dict[v]++
	}
	return dict
}

var str = "I am learning Go and I love Go"

func main() {
	fmt.Println(WordCount(str))
}
