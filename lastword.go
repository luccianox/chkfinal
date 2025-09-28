/*
lastword
Instructions
Write a function LastWord that takes a string and returns its last word followed by a \n.

A word is a section of string delimited by spaces or by the start/end of the string.
Expected function
*/
package main

func LastWord(s string) string {

	j := len(s) - 1

	for j >= 0 && s[j] == ' ' {

		j--
	}
	if j < 0 {

		return "\n"
	}

	i := j

	for i >= 0 && s[i] != ' ' {

		i--
	}
	return s[i+1:j+1] + "\n"
}
