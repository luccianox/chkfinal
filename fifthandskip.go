/*
fifthandskip
Instructions
Write a function FifthAndSkip() that takes a string and returns another string.
The function separates every five characters of the string with a space and removes the sixth one.

If there are spaces in the middle of a word, ignore them and get the first character after
the spaces until you reach a length of 5.
If the string is less than 5 characters return Invalid Input followed by a newline \n.
If the string is empty return a newline \n.
Expected function
*/
package main

import "fmt"

func FifthAndSkip(str string) string {

	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Input\n"
	}
	result := ""
	skips := 0
	count := 0
	// loop to the fifth element
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		}
		if count == 5 {
			result += " "
			count = 0
			skips = 0
		}
		//skip the 6th
		if skips == 0 && count == 0 && len(result) > 0 {
			skips++
			continue
		}
		result += string(str[i])
		count++
	}
	return result + "\n"
}
func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}
