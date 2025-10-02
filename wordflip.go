/*
wordflip
Instructions
Write a function WordFlip() that takes a string as input and returns it in reverse order.

The output should be followed by a newline \n.
If the string is empty, return Invalid Output.
Ignore multiple spaces between words and trim any leading or trailing spaces in the string.
Expected function
*/
package main

import "fmt"

func WordFlip(str string) string {

	if str == "" {
		return "Invalid Output\n"
	}

	// trim spaces
	start, end := 0, len(str)-1
	for start <= end && str[start] == ' ' {
		start++
	}
	for end >= start && str[end] == ' ' {
		end--
	}
	if start > end {
		return "\n"
	}
	//create our word
	word := ""
	words := []string{}

	// loop through str
	for i := start; i <= end; i++ {

		if str[i] != ' ' {
			word += string(str[i])
		} else {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		}
	}
	if word != "" {
		words = append(words, word)
	}
	//reverse words

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {

		words[i], words[j] = words[j], words[i]
	}
	//join them manually

	result := ""
	//loop thru words

	for i := 0; i < len(words); i++ {

		if i > 0 {
			result += " "
		}
		result += words[i]
	}
	return result + "\n"
}
func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}
