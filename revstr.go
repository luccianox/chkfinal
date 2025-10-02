package main

import (
	"fmt"
	"os"
)

func main() {
	// Check that there is exactly 1 argument
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	var words []string
	word := ""

	// Step 1: split manually into words
	for i := 0; i < len(input); i++ {
		ch := input[i]
		if ch != ' ' {
			// build current word
			word += string(ch)
		} else {
			// space = end of word
			words = append(words, word)
			word = ""
		}
	}
	// Add the last word
	if word != "" {
		words = append(words, word)
	}

	// Step 2: reverse and print
	for i := len(words) - 1; i >= 0; i-- {
		fmt.Print(words[i])
		if i > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
