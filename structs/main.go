package main

import "fmt"

func main() {
	fmt.Printf("welcome to Structs")
	Prashanth := user{"Prashanth", 21, 234000.03, "hellokskndkndfkngkdfnkkgnkewdnkjfn"}
	//fmt.Println("the user", Prashanth)
	fmt.Printf("the value of  a %+v", Prashanth)
	fmt.Printf("the value of %+v", Prashanth.Salary)

	// for user := range Prashanth {
	// 	fmt.Printf("the value of %v", user)
	// }
}

type user struct {
	Name    string
	Age     int
	Salary  float32
	Address string
}
