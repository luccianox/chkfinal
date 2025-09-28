/*
lastword
Instructions
Write a function LastWord that takes a string and returns its last word followed by a \n.

A word is a section of string delimited by spaces or by the start/end of the string.
Expected function
*/
package main

func LastWord(s string) string {

	end := len(s) - 1

	for end >= 0 && s[end] == ' ' {

		end--
	}
	if end < 0 {

		return "\n"
	}
	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}
	return s[start+1:end+1] + "\n"
}
