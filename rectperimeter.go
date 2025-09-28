/*
rectperimeter
Instructions
Write a function that takes two int's as arguments, representing the length of width and height
of a rectangle and returning the perimeter of the rectangle.

If one of the arguments is negative it should return -1.
Expected function
*/
package main

func RectPerimeter(w, h int) int {

	if w < 0 || h < 0 {

		return -1
	}
	return 2 * (w + h)
}
