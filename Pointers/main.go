package main

import "fmt"

func main() {
	fmt.Println("welcome to the New Chapter Pointer")

	var ptr *int

	fmt.Println("the value of the Ptr", ptr)

	myNumber := 23

	Myptr := &myNumber
	fmt.Println("the value of the Ptr", Myptr)
	fmt.Println("the value of the Ptr", *Myptr)

	*Myptr = *Myptr * 2
	fmt.Println("the New value of  the PTR", myNumber)

}
