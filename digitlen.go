/*
digitlen
Instructions
Write a function DigitLen() that takes two integers as arguments and returns the times

	the first int can be divided by the second until it reaches zero.

The second int must be between 2 and 36. If not, the function returns -1.
If the first int is negative, reverse the sign and count the digits.
Expected function
*/
package main

import "fmt"

func DigitLen(n, base int) int {
	count := 0
	if base < 2 || base > 36 {
		return -1
	}

	if n < 0 {

		n = -n
	}
	if n == 0 {

		return 1
	}
	for n > 0 {

		n = n / base
		count++
	}
	return count
}
func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}
