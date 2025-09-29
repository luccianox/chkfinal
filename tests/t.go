package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) != 3 {

		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	result := ""

	for _, a := range s1 {
		found := false
		for _, b := range s2 {

			if a == b {

				found = true
				break
			}
		}
		if found {
			dup := false
			for _, c := range result {

				if c == a {
					dup = true
					break
				}
			}
			if !dup {
				result += string(a)
			}
		}
	}
	for _, ch := range result {

		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
