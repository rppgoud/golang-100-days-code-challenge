package main

import "fmt"

func main() {
	fmt.Println("hello world")

	var sampleNo uint = 986687687876
	fmt.Println("we are priting the types", sampleNo)
	fmt.Printf("the is %T \n", sampleNo)

	var islogged bool = false
	fmt.Println("we are priting the types", !islogged)
	fmt.Printf("the is %T \n", islogged)

	var fName string = "prashanth"
	var lName string = "Pavi"
	fmt.Println(fName + lName)

	var smallVal float64 = 234.2321323123124
	fmt.Println("we are priting the types", smallVal)
	fmt.Printf("the is %T \n", smallVal)

	var smallValu float64
	fmt.Println("we are priting the types", smallValu)
	fmt.Printf("the is %T \n", smallValu)

	var sName string
	fmt.Println("we are priting the types", sName)
	fmt.Printf("the is %T \n", sName)

}
