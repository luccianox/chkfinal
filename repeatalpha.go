/*
repeatalpha
Instructions
Write a function called RepeatAlpha that takes a string and displays it repeating each
alphabetical character as many times as its alphabetical index.

'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc...

Expected Function
*/
package main

import "fmt"

func RepeatAlpha(s string) string {

	result := ""

	//loop through s

	for i := 0; i < len(s); i++ {

		index := int(s[i] - 'a' + 1)

		if s[i] >= 'a' && s[i] <= 'z' {

			for j := 0; j < index; j++ {

				result += string(s[i])
			}
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			index := int(s[i] - 'A' + 1)
			for i := 0; i < index; i++ {

				result += string(s[i])
			}
		} else {
			result += string(s[i])
		}
	}
	return result
}
func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
