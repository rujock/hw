package main

import "fmt"

/*
Sample input/output:
The program prints the perimeter and the square of a rectangle given the rectangle sides.
Enter the length and the breadth of the rectangle:
4 5.5
Perimeter: 19
Square: 22
*/
func main() {
	fmt.Println("The program prints the perimeter and the square of a rectangle given the rectangle sides.")
	fmt.Println("Enter the length and the breadth of the rectangle:")
	var x, y float64
	fmt.Scan(&x, &y)
	fmt.Println("Perimeter:", (x+y)*2)
	fmt.Println("Square:", x*y)
}
