/*
zipstring
Instructions
Write a function that takes a string and returns a new string that replaces every
character with the number of duplicates and the character itself, deleting the extra duplications.

The letters are from the latin alphabet list only. Any other character, symbols, shall not be tested.
Expected function
*/
package main

import (
	"fmt"
	"strconv"
)

func ZipString(s string) string {

	count := 1
	result := ""
	//loop through s
	//count duplicates
	//remove the duplicates
	//develop new string

	// loop through s
	runes := []rune(s)
	for i := 1; i < len(s); i++ {
		if runes[i] == runes[i-1] {
			count++
		} else {
			result += strconv.Itoa(count) + string(rune(runes[i-1]))
			count = 1
		}
	}
	return result + "\n"
}
func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
