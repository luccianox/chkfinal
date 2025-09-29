/*
weareunique
Instructions
Write a function that takes two strings's and returns the number of characters that are not included in both,

	without repeating characters.

If there is no unique characters return 0.
If both strings are empty return -1.
Expected function
*/
package main

func inSlice(s []rune, r rune) bool {

	for _, k := range s {

		if k == r {
			return true
		}
	}
	return false
}
func WeAreUnique(str1, str2 string) int {

	if len(str1) == 0 && len(str2) == 0 {

		return -1
	}

	r1 := []rune(str1)
	r2 := []rune(str2)
	used := []int{}
	count := 0
	//loop thru str1
	for _, ch := range str1 {

		if !inSlice(r2, ch) {

			count++
			used = append(used, int(ch))
		}
	}
	//loop thru str2
	for _, ch := range str2 {

		if !inSlice(r1, ch) {
			count++
			used = append(used, int(ch))
		}
	}
	return count
}
