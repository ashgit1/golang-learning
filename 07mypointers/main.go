package main

import "fmt"

func main() {

	fmt.Println("Welcome to a class on pointers")

	var ptr *int
	fmt.Println("Value of pointers is ", ptr)

	myNumber := 23
	var ptr2 = &myNumber
	fmt.Println("Value of the pointer is ", ptr2)
	fmt.Println("Actual Value of the pointer: ", *ptr2)

	*ptr2 = *ptr2 + 2
	fmt.Println("New Value of the pointer: ", myNumber)

}
