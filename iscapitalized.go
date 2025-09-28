/*
iscapitalized
Instructions
Write a function IsCapitalized() that takes a string as an argument and returns true if each word

	in the string begins with either an uppercase letter or a non-alphabetic character.

If any of the words begin with a lowercase letter return false.
If the string is empty return false.
Expected function
*/
package main

func IsCapitalized(s string) bool {

	//empty string
	if len(s) == 0 {

		return false
	}
	//loop thru s

	for i := 0; i < len(s); i++ {

		// first condition
		if i == 0 || s[i-1] == ' ' {
			//second conditon
			if s[i] >= 'a' && s[i] <= 'z' {
				return false
			}
		}
	}
	return true
}
