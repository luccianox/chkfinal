/*
hashcode
Instructions
Write a function called HashCode() that takes a string as an argument and returns a new hashed string.

The hash equation is computed as follows:
(ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range

	of 0 to 127.

If the resulting character is unprintable add 33 to it.
Expected function
*/
package main

func HashCode(dec string) string {
	result := ""
	n := len(dec)

	for _, ch := range dec {

		hashed := (int(ch) + n) % 127
		if hashed < 33 {

			hashed += 33
		}
		result += string(rune(hashed))
	}
	return result
}
