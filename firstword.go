/*
firstword
Instructions
Write a function that takes a string and return a string containing its first word, followed by a newline ('\n').

A word is a sequence of characters delimited by spaces or by the start/end of the argument.
Expected Function
*/
package main

func FirstWord(s string) string {

	end := len(s)

	for i := 0; i < len(s); i++ {

		if s[i] == ' ' {

			end = i
			break
		}
	}
	return s[:end] + "\n"
}
