/*revconcatalternate
Instructions
Write a function RevConcatAlternate() that receives two slices of int as arguments
and returns a new slice with alternated values of each slice in reverse order.

The input slices can have different lengths.
The new slice should start with the elements from the largest slice first and when they became equal size slices,
 it should add an element of the first given slice.
If the slices are of equal length, the new slice should start with an element of the first slice.
Note: you can check the examples bellow for more details.

Expected function*/
func RevConcatAlternate(slice1, slice2 []int) []int {

	i := len(slice1) - 1
	j := len(slice2) - 1

	// figure out who goes first

	turn := 1 // slice 1, slice

	if slice2 > slice1 {

		turn = 2
	}

}