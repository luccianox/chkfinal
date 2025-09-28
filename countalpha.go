/*
countalpha
Instructions
Write a function CountAlpha() that takes a string as an argument and returns the number
of alphabetic characters in the string.

Expected functions
*/
package main

func CountAlpha(s string) int {

	count := 0

	for i := 0; i < len(s); i++ {

		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {

			count++
		}
	}
	return count
}
