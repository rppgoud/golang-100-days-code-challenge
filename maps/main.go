package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to the maps")

	Languagues := make(map[string]string)
	Languagues["py"] = "python"
	Languagues["rb"] = "ruby"
	Languagues["cp"] = "cplus"
	Languagues["jv"] = "java"

	fmt.Println("priting the values", Languagues)
	delete(Languagues, "cp")
	fmt.Println("priting the values", Languagues)

	// loops are intersting part
	for key, value := range Languagues {
		fmt.Printf("the key is %v and value %v\n", key, value)
	}
}
