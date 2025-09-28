/*
reversestrcap
Instructions
Write a program that takes one or more arguments and that, for each argument, puts the last letter
of each word in uppercase and the rest in lowercase. It displays the result followed by a newline ('\n').

If there are no argument, the program displays nothing.
*/
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func toLower(r rune) rune {

	if r >= 'A' && r <= 'Z' {

		return r + 32
	}
	return r
}
func toUpper(r rune) rune {

	if r >= 'a' && r <= 'z' {
		return r - 32
	}
	return r
}
func main() {

	s1 := os.Args[1:]

	for _, ch := range s1 {

		for i, v := range ch {

			if i == len(ch)-1 || ch[i+1] == ' ' {

				z01.PrintRune(toUpper(v))
			} else {
				z01.PrintRune(toUpper(v))
			}
		}
	}
	z01.PrintRune('\n')
}
