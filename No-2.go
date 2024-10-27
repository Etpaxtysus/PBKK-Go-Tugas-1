package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseWord(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter at least 3 words: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		words := strings.Fields(input)

		if len(words) < 3 {
			fmt.Println("Please enter at least 3 words.")
			continue
		}

		for i, word := range words {
			words[i] = reverseWord(word)
		}

		reversed := strings.Join(words, " ")
		fmt.Println("Reversed:", reversed)

		break
	}
	
}
