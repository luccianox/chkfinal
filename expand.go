/*expandstr
Instructions
Write a program that takes a string and displays it with exactly three spaces between each word,
with no spaces nor tabs at neither the beginning nor the end.

The string will be followed by a newline ('\n').

A word, in this exercise, is a sequence of visible characters.

If the number of arguments is not 1, or if there are no word, the program displays nothing.*/

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) < 2 {
		return
	}

	s1 := []rune(os.Args[1])
	words := []string{}

	//loop through s1
	for i := 0; i < len(s1); i++ {
		//skip leading spaces
		for i < len(s1) && (s1[i] == ' ' || s1[i] == '\t') {
			i++
		}
		// select word
		word := ""
		for i < len(s1) && (s1[i] != ' ' || s1[i] != '\t') {
			word += string(s1[i])

		}
		if word != "" {

			words = append(words, word)
		}
	}
	//print words
	for k, v := range words {
		for _, ch := range v {
			z01.PrintRune(ch)
		}
		if k < len(words)-1 {
			z01.PrintRune(' ')
			z01.PrintRune(' ')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
