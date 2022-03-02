package main

import "fmt"

func main() {
	fmt.Println("hellow welcome to the arrays")

	var fruits [4]string

	fruits[0] = "nana"
	fruits[1] = "appli"
	fruits[2] = "doora"
	fmt.Println("fruits list", fruits)
	fmt.Println("fruits list", len(fruits))

	var vegList = [5]string{"hek", "ddkes", "prasd"}
	fmt.Println("hello the veg list", vegList)
	fmt.Println("hello the veg list", len(vegList))

}
