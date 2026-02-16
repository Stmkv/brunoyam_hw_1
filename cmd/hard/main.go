package main

import (
	"fmt"
	"homework/internal/input"
	"strings"
)

func main() {
	userInput := input.ReadLine()
	counter := make(map[string]int)
	words := processString(userInput)
	for _, v := range words {
		counter[v]++
	}
	fmt.Println(counter)

}

func processString(text string) []string {
	replacer := strings.NewReplacer(
		",", "",
		":", "",
		"?", "",
		"!", "",
		"-", " ",
	)
	lowerWords := make([]string, 0)
	text = replacer.Replace(text)
	words := strings.Split(text, " ")
	for _, word := range words {
		lowerWords = append(lowerWords, strings.ToLower(word))
	}

	return lowerWords
}
