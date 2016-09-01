package main

import (
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	wordToInt := make(map[string]int)
	array := strings.Fields(s)
	for i:=0; i < len(array); i++ {
		wordToInt[array[i]] ++;
	}
	return wordToInt
}

func main() {
	fmt.Println(WordCount("foo sss aaa"))
	fmt.Println(WordCount("foo sss aaa aaa aaa foo"))
}
