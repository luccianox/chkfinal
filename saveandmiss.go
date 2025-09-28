/*
saveandmiss
Instructions
Write a function called SaveAndMiss() that takes a string and an int as an argument.
The function should move through the string in sets determined by the int, saving the first set,
omitting the second, saving the third, and so on, in a 'save' and 'miss'
fashion until the end of the string is reached. Return a string containing the saved characters.

If the int is 0 or a negative number return the original string.

Expected function
*/
package main

func SaveAndMiss(arg string, num int) string {

	if num <= 0 {

		return arg
	}
	runes := []rune(arg)
	result := ""
	save := true

	for i := 0; i < len(runes); i += num {

		end := i + num
		if end > len(runes) {
			end = len(runes)
		}
		if save {
			result += string(runes[i:end])
		}
		save = !save
	}
	return result
}
