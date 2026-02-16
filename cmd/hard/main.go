package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	m := make([]string, 0)
	sort.Strings(m)
	userInput := reader()
	counter := make(map[string]int)
	words := processString(userInput)
	for _, v := range words {
		counter[v]++
	}
	fmt.Println(counter)
}

func reader() (text string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text = scanner.Text()
	return
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
